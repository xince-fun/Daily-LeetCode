package main

import "sort"

func largestPerimeter1(nums []int) int64 {
	// 复制并排序原始数组
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)

	// 计算前缀和
	prefixSum := make([]int, len(tmp))
	prefixSum[0] = tmp[0]
	for i := 1; i < len(tmp); i++ {
		prefixSum[i] = tmp[i] + prefixSum[i-1]
	}

	// 寻找满足条件的子数组索引
	index := len(tmp)
	for j := len(tmp) - 1; j >= 0; j-- {
		if j == 0 || prefixSum[j-1] <= tmp[j] {
			index = j
			break
		}
	}

	// 根据索引构建结果数组
	resArr := tmp[:index]
	if len(resArr) >= 3 {
		// 如果结果数组长度至少为3，返回元素和
		sum := 0
		for _, v := range resArr {
			sum += v
		}
		return int64(sum)
	} else {
		// 否则返回-1
		return -1
	}
}

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)

	sum := func(nums []int) int64 {
		var ans int64
		for _, num := range nums {
			ans += int64(num)
		}
		return ans
	}
	// 枚举最大的数字，在 nums[i] 左边的所有数我都要选
	s := sum(nums)

	for i := len(nums) - 1; i > 2; i-- {
		x := int64(nums[i])
		if s-x > x {
			return s
		}
		s -= x
	}
	return -1
}
