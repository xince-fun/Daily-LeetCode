package main

// 上一个更大（小）元素
// 下一个更大（小）元素

func dailyTemperatures1(temperatures []int) []int {
	n := len(temperatures)
	st := []int{}
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		for len(st) > 0 && t >= temperatures[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return ans
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	st := []int{}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		t := temperatures[i]
		// 题目要求什么，while 循环这里就写什么,这里题目要求的是大一个的温度
		for len(st) > 0 && t > temperatures[st[len(st)-1]] {
			x := st[len(st)-1]
			ans[x] = i - x
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
