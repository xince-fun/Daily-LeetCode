package main

// 维护最大最小值

func continuousSubarrays(nums []int) int64 {
	ans, left := int64(0), 0
	q, p := []int{}, []int{}

	for right, num := range nums {
		// q 维护最大值
		for len(q) > 0 && nums[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}
		// p 维护最小值
		for len(p) > 0 && nums[p[len(p)-1]] >= num {
			p = p[:len(p)-1]
		}
		q = append(q, right)
		p = append(p, right)

		for nums[q[0]]-nums[p[0]] > 2 {
			if q[0] == left {
				q = q[1:]
			}
			if p[0] == left {
				p = p[1:]
			}
			left++
		}
		ans += int64(right - left + 1)
	}
	return ans
}
