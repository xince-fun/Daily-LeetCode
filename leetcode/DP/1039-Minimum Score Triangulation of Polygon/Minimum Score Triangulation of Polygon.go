package main

import "math"

// 定义 从i到j 表示沿着边从顶点i顺时针到顶点j，再加上直接从j到i的这条边所组成的多边形

// 子问题： 计算从i到j的最低得分
// 下一个子问题： 计算从i到k的最低得分 计算从k到j的最低得分

// dfs(i,j) = max {dfs(i,k)+dfs(k,j)+v[i]*v[j]*v[k]}

// dfs(i,i+1)=0

func minScoreTriangulation1(values []int) int {
	n := len(values)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i+1 == j {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res = math.MaxInt
		for k := i + 1; k < j; k++ {
			res = min(res, dfs(i, k)+dfs(k, j)+values[i]*values[j]*values[k])
		}
		return res
	}
	return dfs(0, n-1)
}

func minScoreTriangulation(values []int) int {
	n := len(values)

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}

	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			res := math.MaxInt
			for k := i + 1; k < j; k++ {
				res = min(res, f[i][k]+f[k][j]+values[i]*values[j]*values[k])
			}
			f[i][j] = res
		}
	}
	return f[0][n-1]
}
