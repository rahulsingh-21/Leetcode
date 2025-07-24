package main

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}

	inorderArr := inorder(root)
	return inorderArr[k-1]
}

func inorder(root *TreeNode) []int {
	var nums []int
	if root == nil {
		return nums
	}

	nums = append(nums, inorder(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, inorder(root.Right)...)

	return nums
}
