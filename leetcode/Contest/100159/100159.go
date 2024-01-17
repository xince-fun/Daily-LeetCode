package main

func minimumOperationsToMakeEqual1(x int, y int) int {
	if x == y {
		return 0
	}

	visited := make(map[int]bool)
	queue := []struct {
		val int
		ops int
	}{{val: x, ops: 0}}
	visited[x] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// 对当前数字执行四种可能的操作
		operations := []int{
			node.val - 1,
			node.val + 1,
		}
		if node.val%5 == 0 {
			operations = append(operations, node.val/5)
		}
		if node.val%11 == 0 {
			operations = append(operations, node.val/11)
		}

		for _, nextVal := range operations {
			if nextVal == y {
				// 如果达到目标值，返回操作次数
				return node.ops + 1
			}
			if nextVal > 0 && !visited[nextVal] {
				// 如果下一个值有效且未被访问过，将其添加到队列中
				visited[nextVal] = true
				queue = append(queue, struct {
					val int
					ops int
				}{val: nextVal, ops: node.ops + 1})
			}
		}
	}

	return -1 // 这种情况理论上不会发生，因为总是有方法使两数相等
}

// BFS

// 从 v 向 v/11 连边
// 从 v 向 v/5 连边
// 从 v 向 v-1 连边
// 从 v 向 v+1 连边

// 最多是 x - y
// x + x - y + 1

// 记忆化搜索

func minimumOperationsToMakeEqual2(x, y int) int {
	if x <= y {
		return y - x
	}
	ans := x - y
	vis := make([]bool, x+ans+1)
	q := []int{}
	step := 0
	add := func(v int) {
		if v < y {
			ans = min(ans, step+1+y-v)
		} else if !vis[v] {
			vis[v] = true
			q = append(q, v)
		}
	}

	add(x)
	for {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == y {
				return min(ans, step)
			}
			if v%11 == 0 {
				add(v / 11)
			}
			if v%5 == 0 {
				add(v / 5)
			}
			add(v - 1)
			add(v + 1)
		}
		step++
	}
}

func minimumOperationsToMakeEqual(x, y int) int {
	memo := map[int]int{}
	var dfs func(int) int
	dfs = func(x int) int {
		if x <= y {
			return y - x
		}
		if v, ok := memo[x]; ok {
			return v
		}

		res := min(x-y, dfs(x/11)+x%11+1, dfs(x/11+1)-x%11+1, dfs(x/5)+x%5+1, dfs(x/5-x%5+1))
		memo[x] = res
		return res
	}
	return dfs(x)
}
