package main

func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	store := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		store[i] = make([]int, len(nums))
	}

	var rangeBest func(i, j int)
	rangeBest = func(i, j int) {
		m := 0
		// k取值在(i,j)开区间
		for k := i + 1; k < j; k++ {
			left := store[i][k]
			right := store[k][j]
			a := left + nums[i]*nums[k]*nums[j] + right
			if a > m {
				m = a
			}
		}
		store[i][j] = m
	}

	for n := 2; n < len(nums); n++ {
		for i := 0; i < len(nums)-n; i++ {
			rangeBest(i, i+n)
		}
	}
	return store[0][len(nums)-1]
}
