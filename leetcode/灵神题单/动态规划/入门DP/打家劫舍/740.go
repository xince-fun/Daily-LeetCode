package main

// [2,3,4]

// [2,2,3,3,3,4]
// 选不选的问题

// 当前操作？ 枚举 数x 选不选
// 子问题？ 当前选择数后，最大的点数
// 下一个问题？ 分类讨论：
// 不选： 从前 m-1 个数得到的最大点数
//  选： 从前 m-2 个数得到的最大点数

// dfs(i) = max(dfs(i-1), dfs(i-2) + nums[i]*len(n))
// dfs(0) = 0

func deleteAndEarn1(nums []int) int {
	mx := 0
	m := map[int]int{}
	for _, x := range nums {
		mx = max(mx, x)
		m[x]++
	}

	memo := make([]int, mx+1)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		res := max(dfs(i-1), dfs(i-2)+i*m[i])
		*p = res
		return res
	}
	return dfs(mx)
}

func deleteAndEarn2(nums []int) int {
	mx := 0
	m := map[int]int{}
	for _, x := range nums {
		mx = max(mx, x)
		m[x]++
	}

	f := make([]int, mx+1)

	for i := 2; i <= mx; i++ {
		f[i] = max(f[i-1], f[i-2]+i*m[i])
	}
	return f[mx]
}

func deleteAndEarn(nums []int) int {
	mx := 0
	m := map[int]int{}
	for _, x := range nums {
		mx = max(mx, x)
		m[x]++
	}

	f0, f1 := 0, m[1]

	for i := 2; i <= mx; i++ {
		f0, f1 = f1, max(f1, f0+i*m[i])
	}
	return f1
}
