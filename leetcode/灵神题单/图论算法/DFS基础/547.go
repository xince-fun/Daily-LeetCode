package main

func findCircleNum(isConnected [][]int) (ans int) {
	n := len(isConnected)
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !vis[to] {
				dfs(to)
			}
		}
	}

	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return
}
