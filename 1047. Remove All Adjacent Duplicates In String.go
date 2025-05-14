package main

func removeDuplicatesII(s string) string {
	var arr []byte
	var ans string
	arr = append(arr, s[0])
	for i := 1; i < len(s); i++ {
		if len(arr) > 0 && s[i] == arr[len(arr)-1] {
			arr = arr[:len(arr)-1]
		} else {
			arr = append(arr, s[i])
		}

	}
	for _, val := range arr {
		ans += string(val)
	}
	return ans
}
