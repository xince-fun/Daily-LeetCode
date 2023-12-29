package main

func trap1(height []int) int {
	ans, n := 0, len(height)
	// 两个数组
	// 第一个数组 存储前缀最大值， 第二个数组 存储后缀最大值
	preMax, sufMax := make([]int, n), make([]int, n)
	preMax[0] = height[0]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}
	sufMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}

	for i := 0; i < n; i++ {
		ans += min(preMax[i], sufMax[i]) - height[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func trap(height []int) int {
	ans, n := 0, len(height)
	left, right := 0, n-1
	preMax, sufMax := 0, 0
	for left <= right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])
		if preMax < sufMax {
			ans += preMax - height[left]
			left++
		} else {
			ans += sufMax - height[right]
			right--
		}
	}
	return ans
}

func trap2(height []int) (ans int) {
	n := len(height)

	preMax, sufMax := make([]int, n), make([]int, n)
	preMax[0] = height[0]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}
	sufMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}
	for i := 0; i < n; i++ {
		ans += min(preMax[i], sufMax[i]) - height[i]
	}
	return
}

func trap3(height []int) (ans int) {
	n := len(height)

	left, right := 0, n-1
	preMax, sufMax := 0, 0

	for left <= right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])
		if preMax < sufMax {
			ans += preMax - height[left]
			left++
		} else {
			ans += sufMax - height[right]
			right--
		}
	}
	return
}
