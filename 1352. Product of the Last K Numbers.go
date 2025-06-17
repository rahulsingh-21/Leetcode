package main

type ProductOfNumbers struct {
	nums []int
}

func ConstructorVI() ProductOfNumbers {
	return ProductOfNumbers{[]int{1}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.nums = []int{1}
		return
	}
	last := this.nums[len(this.nums)-1]
	this.nums = append(this.nums, num*last)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.nums)

	if k >= n {
		return 0
	}
	return this.nums[n-1] / this.nums[n-1-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
