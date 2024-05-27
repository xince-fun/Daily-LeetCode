package main

func minIncrements1(n int, cost []int) (ans int) {
	var dfs func(int) int
	dfs = func(i int) int {
		if 2*i > n {
			return cost[i-1]
		}
		left := dfs(2 * i)
		right := dfs(2*i + 1)
		ans += max(left, right) - min(left, right)
		return cost[i-1] + max(left, right)
	}
	dfs(1)
	return
}

// 有两条路能走， 左边 2*i+1 右边 2*i+2
func minIncrements(n int, cost []int) (ans int) {
	for i := n / 2; i > 0; i-- {
		left, right := cost[i*2-1], cost[i*2]
		if left > right {
			left, right = right, left
		}
		ans += right - left
		cost[i-1] += right
	}
	return
}
