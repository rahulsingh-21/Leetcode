package main

import "fmt"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := depth(root.Left)
	r := depth(root.Right)
	fmt.Println(l, r)
	if l-r <= 1 && l-r >= -1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func right(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + right(root.Right)
}
