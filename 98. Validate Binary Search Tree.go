package main

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	inorderArr := inorderTraversal(root)
	for i := 0; i < len(inorderArr)-1; i++ {
		if inorderArr[i] >= inorderArr[i+1] {
			return false
		}
	}
	return true
}
