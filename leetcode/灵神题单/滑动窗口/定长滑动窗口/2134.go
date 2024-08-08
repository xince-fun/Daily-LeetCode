package main

// [0,1,0,1,1,0,0]

// [0,1,0,1,1,0,0,0,1,0,1,1,0,0]

// [0,1,1,1,0,0,1,1,0]

// [0,1,1,1,0,0,1,1,0,0,1,1,1,0,0,1,1,0]

func minSwaps1(nums []int) int {
	k := 0
	for _, v := range nums {
		if v == 1 {
			k++
		}
	}
	ans := k
	nums = append(nums, nums...)
	count := 0
	for i, in := range nums {
		if in == 1 {
			count++
		}
		if i < k-1 {
			continue
		}
		ans = min(ans, k-count)
		out := nums[i-k+1]
		if out == 1 {
			count--
		}
	}
	return ans
}

func minSwaps(nums []int) int {
	k := 0
	for _, v := range nums {
		k += v
	}
	cnt1, maxCnt := 0, 0
	nums = append(nums, nums...)
	for i, in := range nums {
		cnt1 += in
		if i >= k {
			cnt1 -= nums[i-k]
			maxCnt = max(maxCnt, cnt1)
		}
	}
	return k - maxCnt
}
