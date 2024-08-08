package main

func numSubarraysWithSum1(nums []int, goal int) (ans int) {
	s := 0
	cnt := map[int]int{}
	cnt[0] = 1
	for _, x := range nums {
		s += x
		ans += cnt[s-goal]
		cnt[s]++
	}
	return ans
}

// 把「恰好」转为「最多」

// 求恰好k个，最用最多k个-最多k-1个

func numSubarraysWithSum2(nums []int, goal int) int {
	cal := func(k int) (ans int) {
		sum := 0
		left := 0
		for right, x := range nums {
			sum += x
			for sum > k && left <= right {
				sum -= nums[left]
				left += 1
			}
			ans += right - left + 1
		}
		return ans
	}
	return cal(goal) - cal(goal-1)
}

func numSubarraysWithSum(nums []int, goal int) (ans int) {
	left1, left2 := 0, 0
	sum1, sum2 := 0, 0
	for right, num := range nums {
		sum1 += num
		for left1 <= right && sum1 > goal {
			sum1 -= nums[left1]
			left1++
		}
		sum2 += num
		for left2 <= right && sum2 >= goal {
			sum2 -= nums[left2]
			left2++
		}
		ans += left2 - left1
	}
	return
}
