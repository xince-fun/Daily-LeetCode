package main

func findOrder1(numCourses int, prerequisites [][]int) []int {
	in, out := make([]int, numCourses), make([][]int, numCourses)
	queue := make([]int, 0)
	count := 0

	for _, pre := range prerequisites {
		out[pre[1]] = append(out[pre[1]], pre[0])
		in[pre[0]]++
	}

	ans := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}

	if len(queue) == 0 {
		return []int{}
	}

	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			count++
			tmp := queue[0]
			ans = append(ans, tmp)
			queue = queue[1:]
			for _, next := range out[tmp] {
				in[next]--
				if in[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
	}
	if count != numCourses {
		return []int{}
	}
	return ans
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	visited := make([]int, numCourses)

	graph := make([]map[int]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make(map[int]bool)
	}
	for _, pre := range prerequisites {
		graph[pre[1]][pre[0]] = true
	}

	ans := []int{}
	var dfs func(int) bool
	dfs = func(i int) bool {
		// 有环
		if visited[i] == 1 {
			return true
		}
		if visited[i] == 2 {
			return false
		}
		visited[i] = 1

		for k := range graph[i] {
			if dfs(k) {
				return true
			}
		}

		visited[i] = 2
		ans = append(ans, i)
		return false
	}

	for i := 0; i < numCourses; i++ {
		if dfs(i) {
			return []int{}
		}
	}
	return reverse(ans)
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
