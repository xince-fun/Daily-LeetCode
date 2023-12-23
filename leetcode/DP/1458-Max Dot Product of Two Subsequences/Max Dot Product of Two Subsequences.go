package main

import "math"

//				dfs(i-1, j-1) + nums1*nums2  选nums1,nums2
// dfs(i, j) =  max(dfs(i-1,j), dfs(i,j-1), dfs(i-1,j-1))

func maxDotProduct1(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	mk := math.MinInt
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		k := nums1[i] * nums2[j]
		mk = max(mk, k)
		memo[i][j] = max(max(dfs(i-1, j), dfs(i, j-1)), dfs(i-1, j-1)+k)
		return memo[i][j]
	}
	res := dfs(n-1, m-1)
	if res > 0 {
		return res
	}
	return mk
}

func maxDotProduct2(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	mk := math.MinInt
	for i, x := range nums1 {
		for j, y := range nums2 {
			k := x * y
			mk = max(mk, k)
			f[i+1][j+1] = max(max(f[i][j+1], f[i+1][j]), f[i][j]+k)
		}
	}
	return f[n][m]
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	m := len(nums2)
	f := make([]int, m+1)
	mk := math.MinInt
	for _, x := range nums1 {
		pre := f[0]
		for j, y := range nums2 {
			temp := f[j+1]
			k := x * y
			mk = max(mk, k)
			f[j+1] = max(max(f[j+1], f[j]), pre+k)
			pre = temp
		}
	}
	if f[m] > 0 {
		return f[m]
	}
	return mk
}
