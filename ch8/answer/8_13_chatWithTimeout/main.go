/*
练习 8.11： 紧接着8.4.4中的mirroredQuery流程，实现一个并发请求url的fetch的变种。当第
一个请求返回时，直接取消其它的请求。
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	name   string
	writer chan<- string // an outgoing message channel
}

var (
	entering    = make(chan client)
	leaving     = make(chan client)
	messages    = make(chan string) // all incoming client messages
	idleTimeout = 30 * time.Second  // idle timeout duration
)

func broadcaster() {
	clients := make(map[client]bool) // set of connected clients

	for {
		select {
		case msg := <-messages:
			// Send the incoming message to all clients
			for cli := range clients {
				select {
				case cli.writer <- msg:
					// Message sent to the client
				default:
					//跳过阻塞的客户端
					// Client is not ready to receive message, skip to the next client
				}
			}

		case cli := <-entering:
			// Add new client to the set of connected clients
			clients[cli] = true
			cli.writer <- "Current clients:"
			for c := range clients {
				cli.writer <- c.name
			}

		case cli := <-leaving:
			// Remove client from the set of connected clients
			delete(clients, cli)
			close(cli.writer)
		}

	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 10) // outgoing client messages
	go clientWriter(conn, ch)

	//who := conn.RemoteAddr().String()
	who := ""
	ch <- "Please enter your name:"
	input := bufio.NewScanner(conn)
	if input.Scan() {
		who = input.Text()
	}
	cli := client{name: who, writer: ch}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	messageCh := make(chan string) // channel to receive client messages

	go func() {
		for input.Scan() {
			messageCh <- input.Text()
		}
	}()

	for {
		select {
		case msg := <-messageCh:
			messages <- who + ": " + msg

		case <-time.After(idleTimeout):
			messages <- who + " has disconnecting"
			leaving <- cli
			time.Sleep(5 * time.Second)
			conn.Close()
			return
		}
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
