/*
练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作
*/
package main

import "fmt"

func main() {
	s := []string{"apple", "apple", "banana", "banana", "banana", "orange", "orange", "orange", "orange"}
	s = eliminateDuplicates(s)
	fmt.Println(s) // 输出 [apple banana orange]
}

func eliminateDuplicates(s []string) []string {
	if len(s) < 2 {
		return s
	}

	i := 0
	for j := 1; j < len(s); j++ {
		if s[j] != s[i] {
			i++
			s[i] = s[j]
		}
	}

	return s[:i+1]
}
