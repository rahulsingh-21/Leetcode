package main

func goodNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}
	val := root.Val
	count := 0

	var dfs func(root *TreeNode, n int)
	dfs = func(root *TreeNode, n int) {
		if root == nil {
			return
		}

		//fmt.Println(root.Val)
		if root.Val >= n {
			n = root.Val
			count++
		}

		dfs(root.Left, n)
		dfs(root.Right, n)
	}

	dfs(root, val)

	return count
}
