package main

// 最后爬了 x = nums[i], 得先爬到 i - x
// dfs(i) = sum(dfs(i-nums[j]))
// dfs(0) = 1 // 爬 0 个台阶的方法数是 1

// dfs(i) = dfs(i-nums[0]) + dfs(i-nums[1]) + dfs(i-nums[2])..
// dfs(0) = 1

func combinationSum41(nums []int, target int) int {
	var dfs func(int) int
	memo := make([]int, target+1)
	for i := range memo {
		memo[i] = -1
	}
	dfs = func(i int) (res int) {
		if i == 0 { // 爬完了
			return 1
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		for _, x := range nums {
			if x <= i {
				res += dfs(i - x)
			}
		}
		*p = res
		return
	}
	return dfs(target)
}

// f[i] = sum(f[i-nums[j]])

func combinationSum4(nums []int, target int) int {
	f := make([]int, target+1)
	f[0] = 1
	for i := 1; i <= target; i++ {
		for _, x := range nums {
			if x <= i {
				f[i] += f[i-x]
			}
		}
	}
	return f[target]
}
