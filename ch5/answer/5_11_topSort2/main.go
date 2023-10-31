/*
练习5.11： 现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图
中的环。
*/
package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"},
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	visiting := make(map[string]bool)

	var visitAll func(string) error
	visitAll = func(item string) error {
		if visiting[item] {
			return fmt.Errorf("cycle detected: %s", item)
		}

		if !seen[item] {
			seen[item] = true
			visiting[item] = true

			for _, prereq := range m[item] {
				err := visitAll(prereq)
				if err != nil {
					return err
				}
			}

			delete(visiting, item)
			order = append(order, item)
		}

		return nil
	}

	for key := range m {
		err := visitAll(key)
		if err != nil {
			return nil, err
		}
	}

	reverse(order)
	return order, nil
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
