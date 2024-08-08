package main

func maxOperations(nums []int) (ans int) {
	num := nums[0] + nums[1]
	for {
		if len(nums) < 2 {
			break
		}
		sum := nums[0] + nums[1]
		if sum != num {
			break
		} else {
			nums = nums[2:]
			ans++
		}
	}
	return
}
