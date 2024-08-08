package main

import "slices"

func minMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)
	vis := make([]bool, n)
	isInitial := make([]bool, n)
	for _, x := range initial {
		isInitial[x] = true
	}

	var nodeId, size int
	var dfs func(int)
	dfs = func(x int) {
		size++
		vis[x] = true
		if nodeId != -2 && isInitial[x] {
			if nodeId < 0 {
				nodeId = x
			} else {
				nodeId = -2
			}
		}
		for y, conn := range graph[x] {
			if conn == 1 && !vis[y] {
				dfs(y)
			}
		}
	}

	ans := -1
	maxSize := 0
	for _, x := range initial {
		if vis[x] {
			continue
		}
		nodeId = -1
		size = 0
		dfs(x)
		if nodeId >= 0 && (size > maxSize || size == maxSize && nodeId < ans) {
			ans = nodeId
			maxSize = size
		}
	}
	if ans < 0 {
		return slices.Min(initial)
	}
	return ans
}
