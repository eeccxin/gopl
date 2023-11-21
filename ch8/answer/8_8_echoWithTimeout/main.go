/*
练习 8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，这样服务器可以在客户
端10秒中没有任何喊话时自动断开连接。
下面的实现有问题
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) //新建goroutines处理连接
	}
}

/*
1.启动一个goroutine,for死循环让他不能断掉
select语句case判断两个channel
一个是10秒后断掉连接
另一个是接收标准输入后发送过来的channel，接收到值后，启动goroutinue输出

2.for循环接收标准输入，接收到后发送给message的channel
*/
func handleConn(c net.Conn) {
	fmt.Println("连接成功")
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	var message = make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(10 * time.Second):
				fmt.Println("超时未输入，断开连接")
				c.Close()
				return
			case mes := <-message:
				wg.Add(1)
				go func(c net.Conn, shout string, delay time.Duration) {
					defer wg.Done()
					//fmt.Fprintln(c, "\t", strings.ToUpper(shout))
					//time.Sleep(delay)
					fmt.Fprintln(c, "\t", shout)
					time.Sleep(delay)
					//fmt.Fprintln(c, "\t", strings.ToLower(shout))
					//ch<-struct{}{}

				}(c, mes, 1*time.Second)

			}
		}
	}()
	for input.Scan() {
		text := input.Text()
		message <- text
	}

	wg.Wait()
	//cw := c.(*net.TCPConn)
	//cw.CloseWrite()

	c.Close()
}

//package main
//
//import (
//	"fmt"
//	"net"
//	"time"
//)
//
//func handleConn(conn net.Conn) {
//	defer conn.Close()
//
//	// 设置超时时间为10秒
//	conn.SetDeadline(time.Now().Add(10 * time.Second))
//
//	for {
//		var buf [512]byte
//
//		select {
//		case <-time.After(10 * time.Second):
//			// 超时时间到达，断开连接
//			fmt.Println("Connection timed out")
//			return
//		default:
//			// 读取客户端发送的数据
//			n, err := conn.Read(buf[:])
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//
//			// 处理客户端发送的数据
//			fmt.Println(string(buf[:n]))
//
//			// 重置超时时间
//			conn.SetDeadline(time.Now().Add(10 * time.Second))
//		}
//	}
//}
//
//func main() {
//	listener, err := net.Listen("tcp", "localhost:8000")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//
//		go handleConn(conn)
//	}
//}
