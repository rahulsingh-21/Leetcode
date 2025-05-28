package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	ans := root
	if root.Val == val {
		return root
	}
	if root.Val > val {
		ans = searchBST(root.Left, val)
	} else {
		ans = searchBST(root.Right, val)
	}

	return ans
}
