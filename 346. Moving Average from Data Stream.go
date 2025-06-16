package main

type MovingAverage struct {
	nums []int
	size int
}

func ConstructorIV(size int) MovingAverage {
	return MovingAverage{[]int{}, size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.nums) == this.size {
		this.nums = this.nums[1:]
	}
	this.nums = append(this.nums, val)

	sum := 0
	for _, n := range this.nums {
		sum += n
	}
	return float64(sum) / float64(len(this.nums))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
