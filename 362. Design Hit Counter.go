package main

import "fmt"

type HitCounter struct {
	count []int
}

func ConstructorVII() HitCounter {
	return HitCounter{[]int{}}
}

func (this *HitCounter) Hit(timestamp int) {
	this.count = append(this.count, timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {
	idx := 0

	for len(this.count) > idx && this.count[idx] <= timestamp-300 {
		fmt.Println(this.count[idx])
		idx++
	}
	return len(this.count) - idx
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
