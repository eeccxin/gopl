package main

/*
练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
*/
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := []byte("Hello, 世界")
	reverse(s)
	fmt.Println(string(s)) // 输出 "界世 ,olleH"
}

// 思路：先每个rune字符反转，再反转整个bytes数组
func reverse(s []byte) {
	// 反转每个UTF-8字符
	start := 0
	for start < len(s) {
		_, size := utf8.DecodeRune(s[start:])
		reverseBytes(s[start : start+size])
		start += size
	}
	// 反转整个切片
	reverseBytes(s)
}

func reverseBytes(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
