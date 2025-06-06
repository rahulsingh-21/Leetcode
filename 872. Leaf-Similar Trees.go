package main

import "slices"

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return slices.Equal(dfs(root1), dfs(root2))
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	return append(dfs(root.Left), dfs(root.Right)...)
}
