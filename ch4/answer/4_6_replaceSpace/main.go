/*
练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考
unicode.IsSpace）替换成一个空格返回
*/
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("Hello      World")
	s = replaceSpaces(s)
	fmt.Println(string(s)) // 输出 "Hello World"
}

func replaceSpaces(s []byte) []byte {
	i := 0
	for j := 0; j < len(s); {
		r, size := utf8.DecodeRune(s[j:])
		// 遍历，当字符是空格时，判断已修改数组的前一个字符是否是空格，分别处理
		//复杂地方在于rune字符，其他的和上一个练习：删除相邻重复字符类似
		if unicode.IsSpace(r) {
			if i == 0 || !unicode.IsSpace(getPreviousRune(s[:i])) {
				s[i] = ' '
				i++
			}
		} else {
			copy(s[i:], s[j:j+size])
			i += size
		}
		j += size
	}

	return s[:i]
}

func getPreviousRune(s []byte) rune {
	_, size := utf8.DecodeLastRune(s)
	return rune(s[len(s)-size])
}

func replaceSpaces1(s []byte) []byte {
	result := s[:0] // 创建一个空切片用于存储修改后的结果
	prevSpace := false

	for _, r := range string(s) {
		if unicode.IsSpace(r) {
			if !prevSpace {
				result = append(result, ' ')
				prevSpace = true
			}
		} else {
			result = append(result, []byte(string(r))...) //多参数通过slice传入的方式
			prevSpace = false
		}
	}

	return result
}
