package main

func canFinish1(numCourses int, prerequisites [][]int) bool {
	in, out := make([]int, numCourses), make([][]int, numCourses)
	queue := make([]int, 0)
	count := 0
	if len(prerequisites) == 0 {
		return true
	}

	for _, pre := range prerequisites {
		out[pre[1]] = append(out[pre[1]], pre[0])
		in[pre[0]]++
	}

	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 没有入度为0的课
	if len(queue) == 0 {
		return false
	}

	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			count++
			tmp := queue[0]
			queue = queue[1:]
			for _, next := range out[tmp] {
				in[next]--
				if in[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
	}
	return count == numCourses
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	visited := make([]int, numCourses)

	graph := make([]map[int]bool, numCourses)
	// 初始化图
	for i := 0; i < numCourses; i++ {
		graph[i] = make(map[int]bool)
	}
	for _, pre := range prerequisites {
		graph[pre[1]][pre[0]] = true
	}

	var dfs func(int) bool
	dfs = func(i int) bool {
		// 遇到环了 直接结束
		if visited[i] == 1 {
			return true
		}
		// 没有遇到环，这个节点访问过了
		if visited[i] == 2 {
			return false
		}
		visited[i] = 1

		for k := range graph[i] {
			if dfs(k) {
				return true
			}
		}
		// i 的所有后继节点都访问完了，没有环
		visited[i] = 2
		return false
	}

	for i := 0; i < numCourses; i++ {
		if dfs(i) {
			return false
		}
	}
	return true
}
