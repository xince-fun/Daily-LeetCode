package main

import (
	"fmt"
)

// 选不选的问题

// 回溯三问
// 当前操作？ 把当前的员工i移出原楼，加入新楼
// 子问题？ 枚举>=i的员工
// 下一个子问题？ 枚举 i+1的员工

func maximumRequests1(n int, requests [][]int) (ans int) {
	m := len(requests)
	path := make([]int, n)

	count := 0
	var dfs func(int)
	dfs = func(i int) {
		if i == m {
			// 判断是否满足所有path为0
			if isValid(path) {
				ans = max(ans, count)
			}
			return
		}
		// 不做操作
		dfs(i + 1)
		// 做操作
		count++
		path[requests[i][0]]--
		path[requests[i][1]]++
		dfs(i + 1)
		// 恢复现场
		count--
		path[requests[i][0]]++
		path[requests[i][1]]--
	}
	dfs(0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isValid(path []int) bool {
	for _, p := range path {
		if p != 0 {
			return false
		}
	}
	return true
}

// 从答案的角度来看

func maximumRequests(n int, requests [][]int) (ans int) {
	m := len(requests)

	count := 0
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

func main() {
	n := 5
	request := [][]int{{0, 0}, {1, 2}, {2, 1}}
	fmt.Println(maximumRequests1(n, request))
}
