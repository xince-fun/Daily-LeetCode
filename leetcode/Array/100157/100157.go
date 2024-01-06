package main

func missingInteger(nums []int) int {
	// 如果数组为空，直接返回 1
	if len(nums) == 0 {
		return 1
	}

	sum := nums[0] // 初始化和为第一个元素
	for i := 1; i < len(nums); i++ {
		// 如果当前元素是前一个元素加一，则继续累加和
		if nums[i] == nums[i-1]+1 {
			sum += nums[i]
		} else {
			break
		}
	}

	x := sum
	for {
		exists := false
		for _, num := range nums {
			if num == x {
				exists = true
				x++
				break
			}
		}
		if !exists {
			return x
		}
	}
}
