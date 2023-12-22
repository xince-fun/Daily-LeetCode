package main

// 最长公共子序列

// 启发思路
// 子序列本质上就是选或不选			不选x，不选y
// 考虑最后一对字母，分别叫x和y ->  不选x，选y
//								选x，不选y
//								选x，选y
// 			|
//        一般化

// 回溯三问
// 当前操作？考虑s[i]和t[j]选或不选
// 子问题？ s的前i个字母和t的前j个字母的LCS长度
// 下一个问题？
// s的前i-1个字母和t的前j-1个字母的LCS长度
// s的前i-1个字母和t的前j个字母的LCS长度
// s的前i个字母和t的前j-1个字母的LCS长度
//				max(dfs(i-1,j),dfs(i,j-1),dfs(i-1,j-1)+1) s[i] = t[j]
// dfs(i,j) =
//				max(dfs(i-1,j),dfs(i,j-1),dfs(i-1,j-1)) s[i] != t[j]

// dfs(i,j) = max(dfs(i-1,j),dfs(i,j-1),dfs(i-1,j-1) + (s[i]=t[j]))

// 不能忽略的问题

// 在 s[i] = t[j]时，需要dfs(i-1,j)和dfs(i,j-1)吗？
// 在 s[i] != t[j]时，需要dfs(i-1,j-1)吗？

// 简化
//			   dfs(i-1,j-1)+1 s[i] = t[j]
// dfs(i,j) =
//			   max(dfs(i-1,j),dfs(i,j-1)) s[i] != t[j]

//

func longestCommonSubsequence1(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if text1[i] == text2[j] {
			memo[i][j] = dfs(i-1, j-1) + 1
		} else {
			memo[i][j] = max(dfs(i-1, j), dfs(i, j-1))
		}
		return memo[i][j]
	}
	return dfs(n-1, m-1)
}

//				f[i-1]f[j-1]+1 s[i] == t[j]
// f[i][j] =
//				max(f[i-1][j], f[i][j-1]) s[i] != t[j]

func longestCommonSubsequence2(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}

	for i, x := range text1 {
		for j, y := range text2 {
			if x == y {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[n][m]
}

// 优化成两个一维数组

func longestCommonSubsequence3(s, t string) int {
	n, m := len(s), len(t)
	f := [2][]int{make([]int, m+1), make([]int, m+1)}
	for i, x := range s {
		for j, y := range t {
			if x == y {
				f[(i+1)%2][j+1] = f[i%2][j] + 1
			} else {
				f[(i+1)%2][j+1] = max(f[i%2][j+1], f[(i+1)%2][j])
			}
		}
	}
	return f[n%2][m]
}

func longestCommonSubsequence(s, t string) int {
	m := len(t)
	f := make([]int, m+1)
	for _, x := range s {
		pre := f[0]
		for j, y := range t {
			temp := f[j+1]
			if x == y {
				f[j+1] = pre + 1
			} else {
				f[j+1] = max(f[j+1], f[j])
			}
			pre = temp
		}
	}
	return f[m]
}
