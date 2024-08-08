package main

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)

	cnt := 0
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		cnt++
		for _, y := range rooms[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	dfs(0)
	return cnt == n
}
