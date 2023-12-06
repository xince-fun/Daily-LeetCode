package main

import "fmt"

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	cnt := make([]int, n)

	for _, t := range trips {
		start, end := t[0], t[1]
		var dfs func(int, int) bool
		dfs = func(x, fa int) bool {
			if x == end {
				cnt[x]++ // 更新经过x的次数
				return true
			}
			for _, y := range g[x] {
				if y != fa && dfs(y, x) {
					cnt[x]++ // x 是 end 的祖先节点，也就在路径上
					return true
				}
			}
			return false
		}
		dfs(start, -1)
	}
	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		notHalf := price[x] * cnt[x] // x 不变
		half := notHalf / 2
		for _, y := range g[x] {
			if y != fa {
				nh, h := dfs(y, x)
				notHalf += min(nh, h)
				half += nh
			}
		}
		return notHalf, half
	}
	return min(dfs(0, -1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 4
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}}
	price := []int{2, 2, 10, 6}
	trips := [][]int{{0, 3}, {2, 1}, {2, 3}}

	fmt.Println(minimumTotalPrice(n, edges, price, trips))
}
