package main

func maxSubarrays(nums []int) (ans int) {
	and := -1 // -1 就是 1111...1, 和任何数 AND 都等于那个数
	for _, x := range nums {
		and &= x
		if and == 0 {
			ans++
			and = -1
		}
	}
	return max(ans, 1)
}
