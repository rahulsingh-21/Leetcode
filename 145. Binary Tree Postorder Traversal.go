package main

func postorderTraversal(root *TreeNode) []int {
	var nums []int
	if root == nil {
		return nums
	}
	nums = append(nums, postorderTraversal(root.Left)...)
	nums = append(nums, postorderTraversal(root.Right)...)
	nums = append(nums, root.Val)
	return nums
}
