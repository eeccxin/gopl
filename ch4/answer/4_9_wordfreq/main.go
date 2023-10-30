/*
练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用
Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordFreq := make(map[string]int)
	totalWords := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		wordFreq[word]++
		totalWords++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "读取输入时发生错误：%v\n", err)
		os.Exit(1)
	}

	fmt.Println("单词\t频数\t频率")
	for word, freq := range wordFreq {
		frequency := float64(freq) / float64(totalWords) * 100
		fmt.Printf("%s\t%d\t%.2f%%\n", word, freq, frequency)
	}
}
