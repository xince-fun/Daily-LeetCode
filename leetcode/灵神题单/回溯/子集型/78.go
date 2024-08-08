package main

func subsets1(nums []int) (ans [][]int) {
	n := len(nums)
	path := []int{}

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(i + 1)
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}
	dfs(0)
	return
}

func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	path := []int{}

	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, append([]int(nil), path...))

		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
