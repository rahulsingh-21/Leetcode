package main

func zigzagLevelOrder(root *TreeNode) [][]int {

	var ans [][]int
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		var level []int
		for i := 0; i < n; i++ {
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

	for i := 0; i < len(ans); i++ {
		//temp := ans[i]
		if i%2 != 0 {
			reverse(ans[i])
		}
	}

	return ans
}

func reverse(n []int) {
	l, r := 0, len(n)-1

	for l < r {
		n[l], n[r] = n[r], n[l]
		l++
		r--
	}
}
