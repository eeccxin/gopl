/*
练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
还是用了3次反转
*/
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s) // 输出 [4 5 1 2 3]
}

// 右旋
func rotate(s []int, k int) {
	reverse(s)
	reverse(s[:k])
	reverse(s[k:])
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
