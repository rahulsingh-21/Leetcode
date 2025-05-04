package main

func coloredCells(n int) int64 {
	var ans int64
	if n == 1 {
		return 1
	}
	ans += int64(n) + int64(n-1)
	temp := ans
	for temp > 0 {
		x := temp - 2
		if x < 0 {
			break
		}
		//fmt.Println(ans, x)
		ans = ans + (2 * x)
		temp = temp - 2
		//fmt.Println(ans, x)
	}
	//fmt.Println(ans)
	return ans
}
