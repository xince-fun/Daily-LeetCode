package main

import "math"

type pair struct {
	day   int
	price int
}

type StockSpanner struct {
	count int
	stack []pair
}

func Constructor() StockSpanner {
	return StockSpanner{
		count: -1,
		stack: []pair{{-1, math.MaxInt}},
	}
}

// price 的上一个更大元素
func (this *StockSpanner) Next(price int) int {
	for price >= this.stack[len(this.stack)-1].price {
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.count++
	this.stack = append(this.stack, pair{this.count, price})
	return this.count - this.stack[len(this.stack)-2].day
}

func main() {

}

// 从左往右看
// 求下一个大的元素
// num > nums[stack[len(stack)-1]]
// 在 for 循环内更新
// x := stack[len(stack)-1]
// stack = stack[:len(stack)-1]
// ans[x] = ...

// 从右往左看
// 求下一个大的元素 x >= nums[stack[len(stack)-1]]
// 在外面去更新答案
