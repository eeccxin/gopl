/*
练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理（输出+inf/-inf)，再编写至少接收1个参数的版本。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(max())               // 输出: -Inf
	fmt.Println(max(1, 2, 3, 4, 5))  // 输出: 5
	fmt.Println(min(10, 8, 6, 4, 2)) // 输出: 2
	fmt.Println(min())               // 输出: +Inf
	fmt.Println(min(7))              // 输出: 7
}

func max(nums ...float64) float64 {
	if len(nums) == 0 {
		return math.Inf(-1)
	}

	maxValue := nums[0]
	for _, num := range nums[1:] {
		if num > maxValue {
			maxValue = num
		}
	}

	return maxValue
}

func min(nums ...float64) float64 {
	if len(nums) == 0 {
		return math.Inf(1)
	}

	minValue := nums[0]
	for _, num := range nums[1:] {
		if num < minValue {
			minValue = num
		}
	}

	return minValue
}
