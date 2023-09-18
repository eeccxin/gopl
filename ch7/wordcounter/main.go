package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordLineCounter struct {
	words int
	lines int
}

func (wc *WordLineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wc.words++
	}

	scanner = bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		wc.lines++
	}

	return len(p), nil
}
func main() {
	text := "Hello, how are you?\nI'm doing great!\n"

	var wc WordLineCounter
	wc.Write([]byte(text))

	fmt.Println("Words:", wc.words)
	fmt.Println("Lines:", wc.lines)
}
