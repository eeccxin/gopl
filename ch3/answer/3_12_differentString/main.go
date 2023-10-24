package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	// 处理正负号
	sign := ""
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		sign = s[:1]
		s = s[1:]
	}

	// 处理小数部分
	dotIndex := strings.Index(s, ".")
	fractionalPart := ""
	if dotIndex != -1 {
		fractionalPart = s[dotIndex:]
		s = s[:dotIndex]
	}

	n := len(s)
	if n <= 3 {
		return sign + s + fractionalPart
	}

	// 计算首个逗号前的数字位数
	firstComma := n % 3
	if firstComma == 0 {
		firstComma = 3
	}

	// 将首个逗号前的数字写入缓冲区
	buf.WriteString(sign + s[:firstComma])

	// 从首个逗号后开始，每隔三位插入逗号
	for i := firstComma; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String() + fractionalPart
}

func main() {
	fmt.Println(comma("1234567890.123"))  // 输出: 1,234,567,890.123
	fmt.Println(comma("-9876543210.987")) // 输出: -9,876,543,210.987
	fmt.Println(comma("+1234.5678"))      // 输出: +1,234.5678
}
