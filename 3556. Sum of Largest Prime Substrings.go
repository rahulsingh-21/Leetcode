package main

import (
	"sort"
	"strconv"
)

func sumOfLargestPrimes(s string) int64 {
	ans := map[int]struct{}{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			temp := s[i:j]
			n, _ := strconv.Atoi(temp)
			if checkPrime(n) {
				ans[n] = struct{}{}
			}
		}
	}

	var set []int

	for key, _ := range ans {
		set = append(set, key)
	}
	sort.Slice(set, func(i, j int) bool {
		return set[i] > set[j]
	})
	//fmt.Println(set)
	var sum int64
	for i := 0; i < len(set); i++ {
		if i == 3 {
			return sum
		}
		sum += int64(set[i])
	}
	return sum
}

func checkPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
