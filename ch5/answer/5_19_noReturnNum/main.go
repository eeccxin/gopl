/*
练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
*/
package main

import "fmt"

func main() {
	result := noReturn()
	fmt.Println("Result:", result)
}

func noReturn() (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = 42 // 设置非零值
		}
	}()

	panic("Something went wrong!") // 触发 panic

	// 这里没有 return 语句
}
