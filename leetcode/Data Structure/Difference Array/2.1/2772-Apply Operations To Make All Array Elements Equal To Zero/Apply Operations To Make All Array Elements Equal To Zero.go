package main

func checkArray(nums []int, k int) bool {
	n := len(nums)
	d := make([]int, n+1)
	sum := 0
	for i, x := range nums {
		sum += d[i]
		x += sum
		if x == 0 {
			continue
		}
		if x < 0 || i+k > n {
			return false
		}
		sum -= x //
		d[i+k] += x
	}
	return true
}

func checkArray1(nums []int, k int) bool {
	n := len(nums)
	d := make([]int, n+1)
	d[0] = nums[0]
	for i := 1; i < n; i++ {
		d[i] = nums[i] - nums[i-1]
	}

	s := 0
	for i := 0; i < n-k+1; i++ {
		m := d[i]
		if m < 0 {
			return false
		}
		d[i] -= m
		d[i+k] += m
	}
	for i := 0; i < n; i++ {
		s += d[i]
	}
	return s == 0
}
