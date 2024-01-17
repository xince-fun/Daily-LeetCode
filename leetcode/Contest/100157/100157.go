package main

func missingInteger(nums []int) int {
	sum := nums[0] // 初始化和为第一个元素
	for i := 1; i < len(nums); i++ {
		// 如果当前元素是前一个元素加一，则继续累加和
		if nums[i] == nums[i-1]+1 {
			sum += nums[i]
		} else {
			break
		}
	}

	has := map[int]bool{}
	for _, x := range nums {
		has[x] = true
	}
	for has[sum] {
		sum++
	}
	return sum
}
