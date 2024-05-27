package main

func distinctDifferenceArray(nums []int) []int {
	m := map[int]int{}
	pre := make([]int, len(nums))
	ans := make([]int, len(nums))
	for i, num := range nums {
		m[num]++
		pre[i] = len(m)
	}
	for i, num := range nums {
		m[num]--
		if m[num] == 0 {
			delete(m, num)
		}
		ans[i] = pre[i] - len(m)
	}
	return ans
}
