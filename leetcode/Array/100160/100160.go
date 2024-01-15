package main

func countSetBitsAtMultiplesOfX(n int64, x int) int64 {
	count := int64(0)
	for i := x; i <= 64; i += x {
		completeCycles := (n + 1) / (1 << i)
		count += completeCycles * (1 << (i - 1))

		remaining := (n + 1) % (1 << i)
		if remaining > (1 << (i - 1)) {
			count += remaining - (1 << (i - 1))
		}
	}
	return count
}

func findMaximumNumber(k int64, x int) int64 {
	left, right := int64(1), int64(2)
	for countSetBitsAtMultiplesOfX(right, x) <= k {
		right *= 2
	}

	for left < right {
		mid := (left + right) / 2
		if countSetBitsAtMultiplesOfX(mid, x) <= k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1

}
