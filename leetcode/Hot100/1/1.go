package main

// O(n^2)
func twoSum1(nums []int, target int) (ans []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans = []int{i, j}
				return
			}
		}
	}
	return ans
}

// O(n)
func twoSum(nums []int, target int) (ans []int) {
	m := map[int]int{}
	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}
