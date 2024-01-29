package main

func minOrAfterOperations(nums []int, k int) int {
	left, right := 0, 0
	for _, num := range nums {
		right |= num
	}
	check := func(mid int) bool {
		ops := 0
		for i := 0; i < len(nums); i++ {
			current := nums[i]
			for j := i + 1; j < len(nums) && (current|mid) != mid; j++ {
				current &= nums[j]
				if (current | mid) == mid {
					ops += j - i
					i = j - 1
					break
				}
			}
		}
		return ops <= k
	}
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {

}
