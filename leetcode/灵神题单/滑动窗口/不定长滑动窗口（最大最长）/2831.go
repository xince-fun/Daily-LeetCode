package main

func longestEqualSubarray1(nums []int, k int) (ans int) {
	posLists := make([][]int, len(nums)+1)
	for i, x := range nums {
		posLists[x] = append(posLists[x], i-len(posLists[x]))
	}

	for _, pos := range posLists {
		if len(pos) <= ans {
			continue
		}
		left := 0
		for right, p := range pos {
			for p-pos[left] > k {
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}

func longestEqualSubarray(nums []int, k int) (ans int) {
	posLists := make([][]int, len(nums)+1)
	for i, x := range nums {
		posLists[x] = append(posLists[x], i)
	}

	for _, pos := range posLists {
		if len(pos) <= ans {
			continue
		}
		left := 0
		// 删除前 pos[right]-pos[left]+1
		// 有 right-left+1个数是相同的
		// 需要删除的事 pos[right]-pos[left]-(right-left)
		for right, p := range pos {
			for p-pos[left]-(right-left) > k {
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}
