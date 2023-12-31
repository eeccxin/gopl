/*
练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等
Unicode中不同的字符类别。
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[string]int)
	categories := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "读取错误：%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			categories["Letter"]++
		} else if unicode.IsDigit(r) {
			categories["Digit"]++
		} else if unicode.IsSpace(r) {
			categories["Space"]++
		} else if unicode.IsPunct(r) {
			categories["Punctuation"]++
		} else {
			categories["Other"]++
		}
		counts[string(r)]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Printf("\ncategory\tcount\n")
	for cat, count := range categories {
		fmt.Printf("%s\t%d\n", cat, count)
	}

	if invalid > 0 {
		fmt.Printf("\n%d 个无效的 UTF-8 字符\n", invalid)
	}
}
