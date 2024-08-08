package main

// 从 n 个数中选 k 个数的组合
// 可以看成是长度固定的子集

// 设 path 长为 m
// 那么还需要选 d = k - m 个数

// 设 当前需要从 [1, i] 这i个数中选数
// 如果 i < d
// 最后必然无法选出 k 个数 不需要继续递归

func combine1(n int, k int) (ans [][]int) {
	path := []int{}

	var dfs func(int)
	dfs = func(i int) {
		d := k - len(path)
		// if i < d {
		// return
		// }
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := i; j > d-1; j-- {
			path = append(path, j)
			dfs(j - 1)
			path = path[:len(path)-1]
		}
	}
	dfs(n)
	return
}

func combine(n int, k int) (ans [][]int) {
	path := []int{}

	var dfs func(int)
	dfs = func(i int) {
		d := k - len(path)
		if i < d {
			return
		}
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		dfs(i - 1)
		path = append(path, i)
		dfs(i - 1)
		path = path[:len(path)-1]
	}
	dfs(n)
	return
}
