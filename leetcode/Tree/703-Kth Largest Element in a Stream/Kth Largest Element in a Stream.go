package main

type minHeap struct {
	k    int
	heap []int
}

func createMinHeap(k int, nums []int) *minHeap {
	heap := &minHeap{k: k, heap: []int{}}
	for _, n := range nums {
		heap.add(n)
	}
	return heap
}

// add 方法
func (this *minHeap) add(num int) {
	if len(this.heap) < this.k {
		this.heap = append(this.heap, num)
		this.up(len(this.heap) - 1)
	} else if num > this.heap[0] {
		this.heap[0] = num
		this.down(0)
	}
}

func (this *minHeap) up(i int) {
	for i > 0 {
		parent := (i - 1) >> 1
		if this.heap[parent] > this.heap[i] {
			this.heap[parent], this.heap[i] = this.heap[i], this.heap[parent]
			i = parent
		} else {
			break
		}
	}
}

func (this *minHeap) down(i int) {
	for 2*i+1 < len(this.heap) {
		child := 2*i + 1 // 左子节点
		if child+1 < len(this.heap) && this.heap[child+1] < this.heap[child] {
			child++
		}
		if this.heap[i] > this.heap[child] {
			this.heap[child], this.heap[i] = this.heap[i], this.heap[child]
			i = child
		} else {
			break
		}
	}
}

type KthLargest struct {
	heap *minHeap
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		heap: createMinHeap(k, nums),
	}
}

func (this *KthLargest) Add(val int) int {
	this.heap.add(val)
	return this.heap.heap[0]
}
