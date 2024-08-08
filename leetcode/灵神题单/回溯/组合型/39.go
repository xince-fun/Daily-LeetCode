package main

func combinationSum(candidates []int, target int) (ans [][]int) {

	path := []int{}
	var dfs func(int, int)
	dfs = func(i, sum int) {

		if sum == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		if i == len(candidates) || sum < 0 {
			return
		}

		dfs(i+1, sum)
		path = append(path, candidates[i])
		dfs(i, sum-candidates[i])
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(0, target)
	return
}
