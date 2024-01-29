package main

// 回溯三问
// 当前操作？ 枚举第i个字符，选或不选：
//	选的话，做拼接，移动到下一个，步数+2；不选的话，移动到下一个，步数+1
// 子问题
//	从下标 >= i的字符串开始看
// 下一个子问题
//  从下标 >= i+1的字符串去看

//
//				min(dfs(i-1, j-1)+1, dfs(i-1, j))	ring[i] == key[j]
// dfs(i, j) =  dfs(i-1, j))					    ring[i] != key[j]

// 枚举选哪个的问题

// 对于 s 和 t

// 选择左边的还是右边的
//				min(dfs(l, j+1) + kl, dfs(r, j+1) + kr)  s[i] != t[j]
// 	dfs(i, j) = dfs(i, j+1)	 s[i] == t[j]

func findRotateSteps1(ring string, key string) int {
	// dp 问题
	n, m := len(ring), len(key)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = -1
		}
	}

	pos := [26]int{}
	// 记录每个字母最后出现的下标
	// pos 就是 left[0]
	for i, b := range ring {
		pos[b-'a'] = i
	}
	// 每个 ring[i] 左边 a-z 的最近下标
	left := make([][26]int, n)
	for i, b := range ring {
		left[i] = pos
		pos[b-'a'] = i
	}
	for i := n - 1; i >= 0; i-- {
		pos[ring[i]-'a'] = i
	}
	// 先算出每个字母首次出现的下标
	// 由于 s 是环形的， 循环结束后 pos 就是 right[n-1]
	right := make([][26]int, n)
	for i := n - 1; i >= 0; i-- {
		right[i] = pos
		pos[ring[i]-'a'] = i
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		// 结束
		if j == m {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if ring[i] == key[j] {
			return dfs(i, j+1)
		}
		l := left[i][key[j]-'a']
		res1 := dfs(l, j+1)
		if l > i {
			res1 += n - l + i
		} else {
			res1 += i - l
		}
		r := right[i][key[j]-'a']
		res2 := dfs(r, j+1)
		if r < i {
			res2 += n - i + r
		} else {
			res2 += r - i
		}
		return min(res1, res2)
	}
	return dfs(0, 0) + m
}

func findRotateSteps(ring string, key string) int {
	n, m := len(ring), len(key)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
}
