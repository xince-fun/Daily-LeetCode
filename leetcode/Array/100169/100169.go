package main

import "sort"

const MOD = 1e9 + 7

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	// 添加外围栅栏
	hFences = append(hFences, 1, m)
	vFences = append(vFences, 1, n)

	// 排序栅栏位置
	sort.Ints(hFences)
	sort.Ints(vFences)

	// 找到最大的正方形边长
	hGaps := getAllGaps(hFences)
	vGaps := getAllGaps(vFences)

	ans := -1
	for k := range hGaps {
		if vGaps[k] == true {
			ans = max(ans, k*k)
		}
	}

	// 返回最大正方形面积的模
	return ans % MOD
}

// 辅助函数：在有序数组中查找元素
func getAllGaps(fences []int) map[int]bool {
	m := make(map[int]bool, 0)
	for i := 0; i < len(fences)-1; i++ {
		for j := i + 1; j < len(fences); j++ {
			gap := fences[j] - fences[i]
			if gap > 0 {
				m[gap] = true
			}
		}
	}
	return m
}
