package main

func countCompleteSubarrays(nums []int) (ans int) {
	hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}
	n := len(hash)

	left := 0
	m := make(map[int]int)
	for _, x := range nums {
		m[x]++
		for len(m) == n {
			y := nums[left]
			m[y]--
			if m[y] == 0 {
				delete(m, y)
			}
			left++
		}
		ans += left
	}
	return
}
