package main

func minimumOperationsToMakeEqual(x int, y int) int {
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
