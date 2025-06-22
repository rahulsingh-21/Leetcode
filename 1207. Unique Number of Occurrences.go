package main

import (
	"fmt"
	"strconv"
)

func uniqueOccurrences(arr []int) bool {
	row := '@'
	for i := 0; i < 60; i++ {
		if i%15 == 0 {
			row++
		}
		fmt.Println(strconv.Itoa(i) + string(row))
	}
	return true
}
