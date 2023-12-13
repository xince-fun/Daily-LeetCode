package main

// 横着算
func trap(height []int) int {
	ans := 0
	stack := []int{}
	for i, h := range height {
		for len(stack) != 0 && h >= height[stack[len(stack)-1]] {
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			bottom := height[x]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			dh := min(height[left], h) - bottom
			ans += dh * (i - left - 1)
		}
		stack = append(stack, i)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
