package main

import (
	"fmt"
	"io"
	"os"
)

type LimitedReader struct {
	r io.Reader
	n int64 //剩余可读字节数
}

func (lr *LimitedReader) Read(p []byte) (n int, err error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > lr.n {
		p = p[:lr.n]
	}
	n, err = lr.r.Read(p)
	if err != nil {
		return 0, err
	}
	lr.n -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func main() {
	file, err := os.Open("ch7/answer/7_5_limitreader/example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	lr := LimitReader(file, 10)
	buf := make([]byte, 5)
	for {
		n, err := lr.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error:", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
