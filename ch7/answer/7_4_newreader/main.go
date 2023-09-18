package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	s     string
	index int //offset
}

// 只是用于读，不做转换
func (r *StringReader) Read(p []byte) (n int, err error) {
	if r.index >= len(r.s) {
		return 0, io.EOF
	}

	n = copy(p, r.s[r.index:])
	r.index += n
	return n, nil
}

func NewReader(s string) *StringReader {
	return &StringReader{s: s}
}

func main() {
	html := "<html><body><h1>Hello, World!</h1></body></html>"
	reader := NewReader(html)

	buf := make([]byte, 10)
	for {
		n, err := reader.Read(buf)
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
