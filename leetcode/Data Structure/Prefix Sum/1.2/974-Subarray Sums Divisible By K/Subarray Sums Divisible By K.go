package main

func subarraysDivByK1(nums []int, k int) (ans int) {
	sum := 0
	m := make(map[int]int)
	m[0%k] = 1

	for _, x := range nums {
		sum = (sum + x) % k
		if sum < 0 {
			sum += k
		}
		if v, ok := m[sum]; ok {
			ans += v
			m[sum]++
		} else {
			m[sum] = 1
		}
	}
	return ans
}

func subarraysDivByK(nums []int, k int) (ans int) {
	sum := 0
	m := make([]int, k)
	m[0] = 1

	for i := 0; i < len(nums); i++ {
		sum = (sum + nums[i]) % k
		if sum < 0 {
			sum += k
		}
		ans += m[sum]
		m[sum]++
	}
	return ans
}
