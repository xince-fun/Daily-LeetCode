package main

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	s := []int{}
	t := []int{}
	for i, num := range nums {
		for len(t) != 0 && num > nums[t[len(t)-1]] {
			ans[t[len(t)-1]] = num
			t = t[:len(t)-1]
		}
		j := len(s) - 1
		for j >= 0 && nums[s[j]] < num {
			j--
		}
		t = append(t, s[j+1:]...)
		s = append(s[:j+1], i)
	}
	return ans
}
