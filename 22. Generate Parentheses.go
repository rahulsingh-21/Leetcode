package main

func generateParenthesis(n int) []string {
	var res []string
	dfsII(n, 0, 0, "", &res)
	return res
}

func dfsII(n, oc, cc int, str string, res *[]string) {
	if oc == n && cc == n {
		*res = append(*res, str)
		return
	}

	if oc < n {
		dfsII(n, oc+1, cc, str+"(", res)
	}
	if oc > cc {
		dfsII(n, oc, cc+1, str+")", res)
	}

}
