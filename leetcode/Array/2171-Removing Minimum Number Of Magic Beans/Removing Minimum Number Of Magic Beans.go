package main

import "sort"

func minimumRemoval(beans []int) (ans int64) {
	sort.Ints(beans)
	sum := 0
	maxM := 0
	for i := 0; i < len(beans); i++ {
		sum += beans[i]
		maxM = max(maxM, (len(beans)-i)*beans[i])
	}
	return int64(sum - maxM)
}

func main() {

}
