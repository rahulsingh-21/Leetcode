package main

import "math"

func maxPathSum(root *TreeNode) int {
	sum := math.MinInt
	var path func(*TreeNode) int

	path = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := max(path(node.Left), 0)
		r := max(path(node.Right), 0)

		sum = max(sum, node.Val+l+r)

		return node.Val + max(l, r)
	}
	path(root)
	return sum
}
