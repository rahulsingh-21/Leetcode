package main

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		var level []int
		for i := 0; i < size; i++ {
			node := q[0]
			level = append(level, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			q = q[1:]
		}
		ans = append(ans, level)
	}
	return ans

}
