/*
练习 8.11： 紧接着8.4.4中的mirroredQuery流程，实现一个并发请求url的fetch的变种。当第
一个请求返回时，直接取消其它的请求。
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func main() {
	urls := []string{
		"http://example.com",
		"http://google.com",
		"http://github.com",
		"http://stackoverflow.com",
	}

	responses := make(chan string)
	cancel := make(chan struct{})
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			select {
			case <-cancel:
				return
			default:
				resp, err := http.Get(url)
				if err != nil {
					responses <- fmt.Sprintf("Error fetching %s: %v", url, err)
					return
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					responses <- fmt.Sprintf("Error reading response body from %s: %v", url, err)
					return
				}

				select {
				case <-cancel:
					return
				default:
					responses <- fmt.Sprintf("URL: %s\n%s", url, body[:20])
				}
			}
		}(url)
	}

	firstResponse := <-responses
	close(cancel) //在第一个返回后，进行close广播
	wg.Wait()

	fmt.Println(firstResponse)
}

func init() {
	// Redirect the standard error to a file
	file, err := os.Create("errors.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	os.Stderr = file
}
