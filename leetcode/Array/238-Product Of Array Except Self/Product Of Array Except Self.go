package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	// ans[i] 现在代表 i 左侧数的乘积
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}
	right := 1
	// right 代表 i 右侧数的乘积
	for i := n - 1; i >= 0; i-- {
		ans[i] = ans[i] * right
		right = right * nums[i]
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
