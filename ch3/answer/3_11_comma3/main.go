package main

import (
	"fmt"
)

func areAnagrams(s1, s2 string) bool {
	// 如果字符串相等，则它们不可能是相互打乱的
	if s1 == s2 {
		return false
	}

	// 如果字符串长度不相等，则它们不可能是相互打乱的
	if len(s1) != len(s2) {
		return false
	}

	// 统计字符出现的次数
	counts := make(map[rune]int)
	for _, ch := range s1 {
		counts[ch]++
	}
	for _, ch := range s2 {
		counts[ch]--
	}

	// 检查字符出现次数是否相等
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(areAnagrams("listen", "silent")) // 输出: true
	fmt.Println(areAnagrams("listen", "listen")) // 输出: false
	fmt.Println(areAnagrams("hello", "world"))   // 输出: false
}
