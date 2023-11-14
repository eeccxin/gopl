/*
练习 8.1： 修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，这个程序可
以同时与多个clock服务器通信，从多服务器中读取时间，并且在一个表格中一次显示所有服
务传回的结果，类似于你在某些办公室里看到的时钟墙。如果你有地理学上分布式的服务器
可以用的话，让这些服务器跑在不同的机器上面；或者在同一台机器上跑多个不同的实例，
这些实例监听不同的端口，假装自己在不同的时区。
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func handleConn(c net.Conn, location string, ch chan<- string) {
	defer c.Close()
	for {
		var buf [512]byte
		n, err := c.Read(buf[:])
		if err != nil {
			ch <- fmt.Sprintf("%s: %s", location, err)
			return
		}
		ch <- fmt.Sprintf("%s: %s", location, string(buf[:n]))
	}
}

func main() {
	locations := flag.String("locations", "", "comma-separated list of locations")
	flag.Parse()

	if *locations == "" {
		fmt.Println("Please provide locations using -locations flag")
		os.Exit(1)
	}

	locs := strings.Split(*locations, ",")

	ch := make(chan string)
	for _, loc := range locs {
		conn, err := net.Dial("tcp", "localhost:"+loc)
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn, loc, ch)
	}

	for msg := range ch {
		fmt.Println(msg)
	}
}
