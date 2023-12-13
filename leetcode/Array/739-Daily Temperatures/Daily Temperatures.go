package main

// 及时去掉无用数据
// 保证栈中数据有序

// 从右到左
// 栈中记录下一个更大的元素的候选项
func dailyTemperatures1(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{}
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		for len(stack) != 0 && t >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}

// 找下一个大的
// 从左到右
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[x] = i - x
		}
		stack = append(stack, i)
	}
	return ans
}
