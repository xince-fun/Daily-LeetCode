package main

import (
	"container/heap"
	"sort"
)

/*
中位数贪心
给定一个数组a，那么取 a 的中位数x
x 到 a 中所有数的距离之和是最小的
*/

func numsGame(nums []int) []int {
	const mod = 1_000_000_007
	ans := make([]int, len(nums))
	left := hp{}
	right := hp{}
	for i, b := range nums {
		b -= i
		if i%2 == 0 {
			heap.Push(&right, -left.pushPop(-b))
			x := right.IntSlice[0] // 中位数
			ans[i] = (right.sum - x + left.sum) % mod
		} else {
			heap.Push(&left, -right.pushPop(b))
			ans[i] = (right.sum + left.sum) % mod
		}
	}
	return ans
}

type hp struct {
	sort.IntSlice
	sum int // 堆中元素之和
}

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)); h.sum += v.(int) }
func (hp) Pop() (_ any)  { return } // 没用到，无需实现

// pushPop 先把 v 入堆，然后弹出并返回堆顶
// 如果 v <= 堆顶，则直接返回 v
func (h *hp) pushPop(v int) int {
	if h.Len() > 0 && v > h.IntSlice[0] {
		h.sum += v - h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}
