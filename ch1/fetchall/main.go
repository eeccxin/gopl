// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	reqNum := 2
	for _, url := range os.Args[1:] {
		go fetch(url, ch, reqNum) // start a goroutine
	}
	for i := 0; i < reqNum*len(os.Args[1:]); i++ {
		fmt.Println(<-ch) // receive from channel ch
	}
	//for range os.Args[1:] {
	//	fmt.Println(<-ch) // receive from channel ch
	//}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, reqNum int) {
	for i := 0; i < reqNum; i++ {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			ch <- fmt.Sprint(err) // send to channel ch
			return
		}

		nbytes, err := io.Copy(io.Discard, resp.Body) // ioutil.Discard 在 Go 1.16 版本中已被弃用
		resp.Body.Close()                             // don't leak resources
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
		secs := time.Since(start).Seconds()
		ch <- fmt.Sprintf("%.2fs  %7d  %s new", secs, nbytes, url)
	}
}

//!-
