package main

func rightSideView(root *TreeNode) []int {
	// nums := []int{}

	if root == nil {
		return []int{}
	}
	q := []*TreeNode{root}
	ans := []int{}

	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			if i == n-1 {
				ans = append(ans, node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			q = q[1:]
		}
	}
	return ans
}
