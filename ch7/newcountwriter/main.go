package main

import (
	"fmt"
	"io"
)

type CountWriter struct {
	writer io.Writer
	count  *int64
}

func (cw CountWriter) Write(p []byte) (n int, err error) {
	n, err = cw.writer.Write(p)
	*cw.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	cw := &CountWriter{writer: w, count: &count}
	return cw, cw.count
}

func main() {
	writer, count := CountingWriter(&myWriter{})
	fmt.Fprint(writer, "Hello, World!")
	fmt.Println(*count) // 输出：13
}

// 自定义的实现了 io.Writer 接口的类型
type myWriter struct{}

func (m myWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
