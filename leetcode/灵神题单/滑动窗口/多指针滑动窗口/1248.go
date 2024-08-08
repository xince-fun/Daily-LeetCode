package main

func numberOfSubarrays1(nums []int, k int) (ans int) {
	left1, left2 := 0, 0
	cnt1, cnt2 := 0, 0
	for right, num := range nums {
		if num%2 == 1 {
			cnt1++
			cnt2++
		}
		for left1 <= right && cnt1 > k {
			if nums[left1]%2 == 1 {
				cnt1--
			}
			left1++
		}
		for left2 <= right && cnt2 >= k {
			if nums[left2]%2 == 1 {
				cnt2--
			}
			left2++
		}
		ans += left2 - left1
	}
	return ans
}

// 如果改成最多k个，里面不符合，所以用right-left+1
func numberOfSubarrays(nums []int, k int) int {
	cal := func(m int) (ans int) {
		left, cnt := 0, 0
		for right, num := range nums {
			cnt += num % 2
			for cnt > m {
				cnt -= nums[left] % 2
				left++
			}
			ans += right - left + 1
		}
		return ans
	}
	return cal(k) - cal(k-1)
}
