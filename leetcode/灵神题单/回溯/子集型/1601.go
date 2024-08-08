package main

func maximumRequests1(n int, requests [][]int) (ans int) {
	m := len(requests)
	count := 0

	isValid := func(p []int) bool {
		for _, x := range p {
			if x != 0 {
				return false
			}
		}
		return true
	}

	path := make([]int, n)

	var dfs func(int)
	dfs = func(i int) {
		if i == m {
			if isValid(path) {
				ans = max(ans, count)
			}
			return
		}
		// 不做操作
		dfs(i + 1)
		count++
		path[requests[i][0]]--
		path[requests[i][1]]++
		dfs(i + 1)
		path[requests[i][0]]++
		path[requests[i][1]]--
	}
	dfs(0)
	return
}

func maximumRequests(n int, requests [][]int) (ans int) {
	m := len(requests)
	count := 0

	isValid := func(p []int) bool {
		for _, x := range p {
			if x != 0 {
				return false
			}
		}
		return true
	}

	path := make([]int, n)

	var dfs func(int)
	dfs = func(i int) {
		if isValid(path) {
			ans = max(ans, count)
		}
		if i == m {
			return
		}
		for j := i; j < m; j++ {
			path[requests[j][0]]--
			path[requests[j][1]]++
			count++
			dfs(j + 1)
			count--
			path[requests[j][0]]++
			path[requests[j][1]]--
		}
	}
	dfs(0)
	return
}
