package main

func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	n := len(graph)

	a := []int{0}
	var dfs func(int)
	dfs = func(x int) {
		if x == n-1 {
			ans = append(ans, append([]int(nil), a...))
			return
		}
		for _, y := range graph[x] {
			a = append(a, y)
			dfs(y)
			a = a[:len(a)-1]
		}
	}
	dfs(0)
	return
}
