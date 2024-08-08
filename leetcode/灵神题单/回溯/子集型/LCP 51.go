package main

import "math"

// 选不选的角度
func perfectMenu1(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
	n := len(cookbooks)

	canChoose := func(ma []int, cb []int) bool {
		for i := range ma {
			if ma[i] < cb[i] {
				return false
			}
		}
		return true
	}

	ans := math.MinInt
	sumX, sumY := 0, 0
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			if sumY >= limit {
				ans = max(ans, sumX)
			}
			return
		}
		// 不选
		dfs(i + 1)
		// 选
		if canChoose(materials, cookbooks[i]) {
			sumX += attribute[i][0]
			sumY += attribute[i][1]
			for j := range materials {
				materials[j] -= cookbooks[i][j]
			}
			dfs(i + 1)
			sumX -= attribute[i][0]
			sumY -= attribute[i][1]
			for j := range materials {
				materials[j] += cookbooks[i][j]
			}
		}
	}
	dfs(0)
	if ans == math.MinInt {
		return -1
	}
	return ans
}

func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
	n := len(cookbooks)

	canChoose := func(ma []int, cb []int) bool {
		for i := range ma {
			if ma[i] < cb[i] {
				return false
			}
		}
		return true
	}

	ans := math.MinInt
	sumX, sumY := 0, 0
	var dfs func(int)

	dfs = func(i int) {
		if sumY >= limit {
			ans = max(ans, sumX)
		}
		for j := i; j < n; j++ {
			if canChoose(materials, cookbooks[j]) {
				sumX += attribute[j][0]
				sumY += attribute[j][1]
				for k := range materials {
					materials[k] -= cookbooks[j][k]
				}
				dfs(j + 1)
				sumX -= attribute[j][0]
				sumY -= attribute[j][1]
				for k := range materials {
					materials[k] += cookbooks[j][k]
				}
			}
		}
	}
	dfs(0)
	if ans == math.MinInt {
		return -1
	}
	return ans
}
