package main

/*
二叉树遍历
*/

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 中序遍历
func inorderTraversal(root *Node) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left)
	fmt.Println(root.Value)
	inorderTraversal(root.Right)
}

// 先序遍历
func preorderTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 后序遍历
func postorderTraversal(root *Node) {
	if root == nil {
		return
	}

	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	fmt.Println(root.Value)
}

func main() {
	// 构建一个二叉树
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	// 执行中序遍历
	fmt.Println("中序遍历：")
	inorderTraversal(root)
	fmt.Println("先序遍历：")
	preorderTraversal(root)
	fmt.Println("后序遍历：")
	postorderTraversal(root)
}
