package main

import "container/heap"

const mod = 1e9 + 7

func getNumberOfBacklogOrders(orders [][]int) (ans int) {
	sellHp := &MinHeap{}
	buyHp := &MaxHeap{}
	heap.Init(sellHp)
	heap.Init(buyHp)

	for _, order := range orders {
		price, amount, orderType := order[0], order[1], order[2]
		if orderType == 0 { // 采购
			if sellHp.Len() > 0 {
				for sellHp.Len() > 0 && amount > 0 {
					o := heap.Pop(sellHp).(pair)
					if o.x <= price {
						if o.y >= amount { // 清空amount, o还没清空
							o.y -= amount
							amount -= amount
							heap.Push(sellHp, o)
						} else {
							amount -= o.y
						}
					} else {
						heap.Push(buyHp, pair{price, amount})
						break
					}
				}
			} else {
				heap.Push(buyHp, pair{price, amount})
			}
		} else {
			if buyHp.Len() > 0 {
				for buyHp.Len() > 0 && amount > 0 {
					o := heap.Pop(buyHp).(pair)
					if o.x >= price {
						if o.y >= amount {
							o.y -= amount
							amount -= amount
							heap.Push(buyHp, o)
						} else {
							amount -= o.y
						}
					} else {
						heap.Push(sellHp, pair{price, amount})
						break
					}
				}
			} else {
				heap.Push(sellHp, pair{price, amount})
			}
		}
	}

	for sellHp.Len() > 0 {
		x := heap.Pop(sellHp).(pair)
		ans += x.y % mod
	}
	for buyHp.Len() > 0 {
		x := heap.Pop(buyHp).(pair)
		ans += x.y % mod
	}
	return
}

type pair struct{ x, y int }

type MaxHeap []pair

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i].x > h[j].x }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type MinHeap []pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i].x < h[j].x }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
