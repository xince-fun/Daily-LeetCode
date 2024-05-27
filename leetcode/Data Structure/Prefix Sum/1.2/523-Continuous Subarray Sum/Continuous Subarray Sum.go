package main

func checkSubarraySum(nums []int, k int) bool {
	sum := 0
	m := make(map[int]int)
	m[0] = -1

	for i := 0; i < len(nums); i++ {
		sum = (sum + nums[i]) % k
		if v, ok := m[sum]; ok {
			if i-v >= 2 {
				return true
			}
		} else {
			m[sum] = -1
		}
	}
	return false
}
