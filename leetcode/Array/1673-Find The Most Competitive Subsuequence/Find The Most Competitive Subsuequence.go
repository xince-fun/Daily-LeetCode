package main

// O(n)
func mostCompetitive(nums []int, k int) []int {
	st := []int{}
	for i, x := range nums {
		for len(st) > 0 && x < st[len(st)-1] && len(st)+len(nums)-i > k {
			st = st[:len(st)-1]
		}
		if len(st) < k {
			st = append(st, x)
		}
	}
	return st
}
