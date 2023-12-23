package main

func findNumberOfLIS1(nums []int) int {
	n := len(nums)

	memo := make([]int, n)
	memo2 := make([]int, n)
	var dfs func(int) (int, int)
	dfs = func(i int) (int, int) {
		subLen := 0
		cnt := 0
		if memo[i] > 0 {
			return memo[i], memo2[i]
		}
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				res, len := dfs(j)
				if res == subLen {
					cnt += len
				} else if res > subLen {
					cnt = len
					subLen = res
				}
			}
		}
		memo[i] = subLen + 1
		if cnt != 0 {
			memo2[i] = cnt
		} else {
			memo2[i] = 1
		}
		return memo[i], memo2[i]
	}
	maxLen, count := 0, 0
	for i := 0; i < n; i++ {
		length, cnt := dfs(i)
		if length > maxLen {
			maxLen = length
			count = cnt
		} else if length == maxLen {
			count += cnt
		}
	}

	return count
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		subLen := f[i][0]
		cnt := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if f[j][0] == subLen {
					cnt += f[j][1]
				} else if f[j][0] > subLen {
					cnt = f[j][1]
					subLen = f[j][0]
				}
			}
		}
		if cnt != 0 {
			f[i][1] = cnt
		} else {
			f[i][1] = 1
		}
		f[i][0] = subLen + 1
	}
	maxLen, count := 0, 0
	for i := 0; i < n; i++ {
		length, cnt := f[i][0], f[i][1]
		if length > maxLen {
			maxLen = length
			count = cnt
		} else if length == maxLen {
			count += cnt
		}
	}
	return count
}
