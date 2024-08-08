package main

import (
	"container/heap"
	"slices"
	"sort"
)

func totalCost(costs []int, k int, candidates int) (ans int64) {
	n := len(costs)
	// 一定可以选到cost中的k个数
	if candidates*2+k > n {
		slices.Sort(costs)
		for _, x := range costs[:k] {
			ans += int64(x)
		}
		return
	}

	pre := hp{costs[:candidates]}
	suf := hp{costs[n-candidates:]}
	heap.Init(&pre)
	heap.Init(&suf)
	for i, j := candidates, n-1-candidates; k > 0; k-- {
		if pre.IntSlice[0] <= suf.IntSlice[0] {
			ans += int64(pre.replace(costs[i]))
			i++
		} else {
			ans += int64(suf.replace(costs[j]))
			j--
		}
	}

	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) replace(v int) int { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(h, 0); return top }
