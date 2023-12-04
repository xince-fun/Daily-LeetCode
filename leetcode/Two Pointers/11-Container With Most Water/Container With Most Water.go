package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		h := min(height[left], height[right])
		ans = max(h*(right-left), ans)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
