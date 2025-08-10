package main

type StockSpanner struct {
	stack []int
}

func Constructo() StockSpanner {
	return StockSpanner{[]int{}}
}

func (this *StockSpanner) Next(price int) int {
	this.stack = append(this.stack, price)
	count := 0
	for i := len(this.stack) - 1; i >= 0; i-- {
		if this.stack[i] <= price {
			count++
		} else {
			return count
		}
	}

	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
