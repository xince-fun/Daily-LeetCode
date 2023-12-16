package main

func rob1(nums []int) int {
	f := func(nums []int) int {
		cur, pre := 0, 0
		for _, num := range nums {
			cur, pre = max(pre+num, cur), cur
		}
		return cur
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(f(nums[:len(nums)-1]), f(nums[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	f := func(nums []int) int {
		f0, f1 := 0, 0
		for _, x := range nums {
			f0, f1 = f1, max(f1, f0+x)
		}
		return f1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(f(nums[:len(nums)-1]), f(nums[1:]))
}

func main() {

}
