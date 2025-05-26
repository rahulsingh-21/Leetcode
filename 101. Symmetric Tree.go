package main

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(n *TreeNode, m *TreeNode) bool {
	if n == nil && m == nil {
		return true
	}
	if n == nil || m == nil {
		return false
	}
	if n.Val != m.Val {
		return false
	}
	return isMirror(n.Left, m.Right) && isMirror(n.Right, m.Left)
}
