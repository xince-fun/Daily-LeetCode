package main

// 差分数组
// d[i] = a[0]   i = 0
// d[i] = a[i] - a[i-1] i >= 1

// 性质1: 从左到累累加d中的元素，可以得到数组a
// 操作等价:
// 1. a的子数组a[i],a[i+1],...a[j] 都加上x
// d[i] 增加 x，d[j+1] 减少 x

func carPooling(trips [][]int, capacity int) bool {
	d := [1001]int{}
	for _, t := range trips {
		d[t[1]] += t[0]
		d[t[2]] -= t[0]
	}
	s := 0
	for _, v := range d {
		s += v
		if s > capacity {
			return false
		}
	}
	return true
}

func main() {
	trips := [][]int{{3, 3, 7}, {2, 1, 5}}
	carPooling(trips, 4)
}
