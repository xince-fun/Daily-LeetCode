package main

import "sort"

// 递增子序列 IS, Increasing Subsequence
// 最长递增子序列 LIS, Longest Increasing Subsequence

// O(n^2) 回溯 -> 记忆化搜索 -> 递推
// O(nlogn) 二分 + 贪心

// 枚举选哪个

// 枚举nums[i]作为LIS的末尾元素，那么需要枚举nums[j]作为LIS的倒数第二个元素
// 其中j < i 且 nums[j] < nums[i]

// 回溯三问
// 子问题？ 以nums[i]结尾的LIS长度
// 当前操作？ 枚举nums[j]
// 下一个子问题？ 以nums[j]结尾的LIS长度

// dfs(i) = max{dfs(j)} + 1, j < i 且 nums[j] < nums[i]

// 超时
func lengthOfLIS1(nums []int) int {
	n := len(nums)

	var dfs func(int) int
	dfs = func(i int) int {
		res := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				res = max(res, dfs(j))
			}
		}
		return res + 1
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i))
	}
	return ans
}

func lengthOfLIS2(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	var dfs func(int) int
	dfs = func(i int) int {
		res := 0
		if memo[i] > 0 {
			return memo[i]
		}
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				res = max(res, dfs(j))
			}
		}
		memo[i] = res + 1
		return memo[i]
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i))
	}
	return ans
}

// f[i] = max{f[j]} + 1, j < i 且 nums[j] < nums[i]

func lengthOfLIS3(nums []int) (ans int) {
	n := len(nums)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] += 1
	}
	for i := 0; i < n; i++ {
		if f[i] > ans {
			ans = f[i]
		}
	}
	return
}

// 进阶技巧： 交换状态与状态值

// f[i] 表示末尾元素为nums[i]的LIS长度
// g[i] 表示长度为i+1的IS的末尾元素的最小值

// 例如 nums = [1,6,7,2,4,5,3]
// g = [1]
// g = [1, 6]
// g = [1, 6, 7]
// g = [1, 2, 7]
// g = [1, 2, 4]
// g = [1, 2, 4, 5]
// g = [1, 2, 3, 5]

func lengthOfLIS(nums []int) (ans int) {
	g := nums[:0]
	for _, num := range nums {
		j := sort.SearchInts(g, num)
		if j == len(g) {
			g = append(g, num)
		} else {
			g[j] = num
		}
	}
	return len(g)
}
