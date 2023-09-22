package main

import (
	"fmt"
	"sort"
)

// IsPalindrome 函数接受一个 sort.Interface 类型的参数 s，并判断序列 s 是否是回文序列
func IsPalindrome(s sort.Interface) bool {
	length := s.Len()
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Less(i, j int) bool {
	// 返回true表示需要交换
	return s[i] < s[j]
}
func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func main() {
	ints := IntSlice{1, 2, 3, 2, 1}
	fmt.Println("IsPalindrome:", IsPalindrome(ints)) // Output: true

	ints2 := IntSlice{1, 2, 3, 4, 5}
	fmt.Println("IsPalindrome:", IsPalindrome(ints2)) // Output: false
}
