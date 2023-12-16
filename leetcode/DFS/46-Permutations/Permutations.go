package main

func permute1(nums []int) (ans [][]int) {
	n := len(nums)
	path := make([]int, n)
	onPath := make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j, on := range onPath {
			if !on {
				path[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}
	dfs(0)
	return
}

// 数组path记录路径上的数（已选的数字）
// 集合s记录剩余未选的数字

// 回溯三问
// 1. 当前操作？ 从s中枚举path[i]要填入的数字x
// 2. 子问题？ 构造排列>=i的部分，剩余未选数字部分集合为s
// 3. 下一个子问题？ 构造排列>=i+1的部分，剩余未选数字部分集合为s-{x}

// 也可以用一个布尔数组onPath记录在path中的数字
// 如果nums[i] 在path中，则onPath[i]为真

// dfs(i,s) -> dfs(i+1, s-{x1})
// dfs(i,s) -> dfs(i+1, s-{x2})
// dfs(i,s) -> dfs(i+1, s-{x3})

func permute2(nums []int) (ans [][]int) {
	n := len(nums)
	path := make([]int, n)
	onPath := make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j, on := range onPath {
			if !on {
				path[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}
	dfs(0)
	return
}
