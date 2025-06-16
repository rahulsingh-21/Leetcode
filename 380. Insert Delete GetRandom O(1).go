package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	set  map[int]int
	nums []int
}

func ConstructorII() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{make(map[int]int), []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}

	this.nums = append(this.nums, val)
	this.set[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.set[val]
	if !ok {
		return false
	}
	last := this.nums[len(this.nums)-1]
	this.nums[idx] = last
	this.set[last] = idx
	this.nums = this.nums[:len(this.nums)-1]

	delete(this.set, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {

	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
