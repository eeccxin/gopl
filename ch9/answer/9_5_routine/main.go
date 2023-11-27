package main

import (
	"fmt"
	"time"
)

/*
练习 9.5: 写一个有两个goroutine的程序，两个goroutine会向两个无buffer channel反复地发送
ping-pong消息。这样的程序每秒可以支持多少次通信？
*/

func ping(ch chan<- string) {
	for {
		ch <- "ping"
	}
}

func pong(ch chan<- string) {
	for {
		ch <- "pong"
	}
}

func main() {
	pingCh := make(chan string)
	pongCh := make(chan string)

	go ping(pingCh)
	go pong(pongCh)

	start := time.Now()

	num := 100000
	for i := 0; i < num; i++ {
		select {
		case msg := <-pingCh:
			fmt.Println(msg)
		case msg := <-pongCh:
			fmt.Println(msg)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
	fmt.Printf("Communications per second: %.2f\n", float64(num)/elapsed.Seconds())
}
