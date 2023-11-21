/*
练习 8.9： 编写一个du工具，每隔一段时间将root目录下的目录大小计算并显示出来。
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var interval int
	flag.IntVar(&interval, "interval", 5, "interval in seconds")
	flag.Parse()

	root := os.Args[0]
	if len(root) == 0 {
		root = "."
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			size, err := calculateDirSize(root)
			if err != nil {
				fmt.Printf("Error calculating directory size: %v\n", err)
			} else {
				fmt.Printf("Directory size: %d bytes\n", size)
			}
		}
	}
}

func calculateDirSize(dirPath string) (int64, error) {
	var size int64

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return nil
	})

	return size, err
}
