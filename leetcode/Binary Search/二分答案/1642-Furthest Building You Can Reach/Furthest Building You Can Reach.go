package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func furthestBuilding1(heights []int, bricks int, ladders int) int {
	// 可以二分 [lef,t right]
	// 二分建筑的下标
	left, right := 0, len(heights)-1
	check := func(d int) bool {
		arr := []int{}
		for i := 1; i <= d; i++ {
			if heights[i] > heights[i-1] {
				arr = append(arr, heights[i]-heights[i-1])
			}
		}
		sort.Ints(arr)
		for i := 0; i < ladders && len(arr) > 0; i++ {
			arr = arr[:len(arr)-1]
		}
		sum := 0
		for _, num := range arr {
			sum += num
		}
		return sum <= bricks
	}
	for left <= right {
		// 循环不变量：
		// left-1 的回答一定是「是」
		// right+1 的回答一定是「否」
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func furthestBuilding2(heights []int, bricks int, ladders int) int {
	return sort.Search(len(heights), func(d int) bool {
		arr := []int{}
		for i := 1; i <= d; i++ {
			if heights[i] > heights[i-1] {
				arr = append(arr, heights[i]-heights[i-1])
			}
		}
		sort.Ints(arr)
		for i := 0; i < ladders && len(arr) > 0; i++ {
			arr = arr[:len(arr)-1]
		}
		sum := 0
		for _, num := range arr {
			sum += num
		}
		return sum > bricks
	}) - 1
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] > h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { (*h) = append(*h, x.(int)) }
func (h *hp) Pop() any          { v := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return v }

// 最大堆
func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := hp{}
	heap.Init(&h)
	for i := 1; i < len(heights); i++ {
		if heights[i]-heights[i-1] > 0 {
			bricks -= heights[i] - heights[i-1]
			heap.Push(&h, heights[i]-heights[i-1])
			if bricks < 0 {
				if ladders == 0 {
					return i - 1
				}
				v := heap.Pop(&h).(int)
				ladders--
				bricks += v
			}
		}
	}
	return len(heights) - 1
}

func main() {
	heights := []int{4, 2, 7, 6, 9, 14, 12}
	bricks := 5
	ladders := 1
	fmt.Println(furthestBuilding(heights, bricks, ladders))
}
