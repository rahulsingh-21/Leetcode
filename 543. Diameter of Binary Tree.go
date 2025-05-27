package main

func diameterOfBinaryTree(root *TreeNode) int {

	maxi := 0
	var height func(*TreeNode) int
	height = func(root *TreeNode) int {

		if root == nil {
			return 0
		}
		l := height(root.Left)
		r := height(root.Right)

		maxi = max(maxi, l+r)

		return 1 + max(l, r)
	}
	return max(maxi, height(root)-1)
}
