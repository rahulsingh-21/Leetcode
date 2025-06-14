package main

func replaceElements(arr []int) []int {

	ans := make([]int, len(arr))
	maxi := -1
	ans[len(arr)-1] = -1

	for i := len(arr) - 2; i >= 0; i-- {
		maxi = max(maxi, arr[i+1])
		ans[i] = maxi
	}
	return ans
}
