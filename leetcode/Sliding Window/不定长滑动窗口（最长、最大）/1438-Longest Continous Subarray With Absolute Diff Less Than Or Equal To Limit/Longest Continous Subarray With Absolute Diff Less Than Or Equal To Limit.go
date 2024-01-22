package main

func longestSubarray(nums []int, limit int) (ans int) {
	left := 0
	var q, p []int

	for right, x := range nums {
		// 维护单调性
		// 维护最大
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}
		// 维护最小
		for len(p) > 0 && nums[p[len(p)-1]] >= x {
			p = p[:len(p)-1]
		}
		q = append(q, right)
		p = append(p, right)

		for len(q) > 0 && len(p) > 0 && nums[q[0]]-nums[p[0]] > limit {
			if nums[left] == nums[q[0]] {
				q = q[1:]
			}
			if nums[left] == nums[p[0]] {
				p = p[1:]
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
