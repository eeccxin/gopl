package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

	n := len(s)
	if n <= 3 {
		return s
	}

	// 计算首个逗号前的数字位数
	firstComma := n % 3
	if firstComma == 0 {
		firstComma = 3
	}

	// 将首个逗号前的数字写入缓冲区
	buf.WriteString(s[:firstComma])

	// 从首个逗号后开始，每隔三位插入逗号
	for i := firstComma; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("1234567890")) // 输出: 1,234,567,890
}
