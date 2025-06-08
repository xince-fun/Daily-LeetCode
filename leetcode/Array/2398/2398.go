package main

import (
	"fmt"
	"sort"
)

// 随着 k的增加， 这个总开销会增大，因此具有单调性
func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	n := len(chargeTimes)
	pre := make([]int, n+1)
	for i, v := range runningCosts {
		pre[i+1] = pre[i] + v
	}
	// 我要找的最大的k，使 f(k) <= budget，是一个先 true，后false的 f(k)
	// 随着 k 的变大， 这里先true后false, 所以 可以判断 f(k+1)，同时判断  f(k+1)的值 > budget, 这样就正好
	// 二分答案
	return sort.Search(n, func(k int) bool {
		k++
		queue := &MyQueue{}
		for i := 0; i < k; i++ {
			queue.Push(chargeTimes[i])
		}
		return queue.Front()+k*pre[k] > int(budget)
	})
}

type MyQueue struct {
	deque []int
}

func (this *MyQueue) Pop(value int) {
	if len(this.deque) != 0 && value == this.deque[0] {
		this.deque = this.deque[1:]
	}
}

func (this *MyQueue) Push(value int) {
	for len(this.deque) != 0 && value > this.deque[len(this.deque)-1] {
		this.deque = this.deque[:len(this.deque)-1]
	}
	this.deque = append(this.deque, value)
}

func (this *MyQueue) Front() int {
	return this.deque[0]
}

func main() {
	chargeTimes := []int{3, 6, 1, 3, 4}
	runningCosts := []int{2, 1, 3, 4, 5}
	budget := int64(25)
	fmt.Println(maximumRobots(chargeTimes, runningCosts, budget))
}
