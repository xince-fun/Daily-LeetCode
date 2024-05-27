package main

func platesBetweenCandles(s string, queries [][]int) []int {
	pre := make([]int, len(s)+1)
	left := make([]int, len(s)) // 表示 i 左侧最近蜡烛的位置
	p := -1
	for i, b := range s {
		pre[i+1] = pre[i]
		if b == '*' {
			pre[i+1]++
		} else {
			p = i
		}
		left[i] = p
	}

	right := make([]int, len(s))
	for i, p := len(s)-1, len(s); i >= 0; i-- {
		if s[i] == '|' {
			p = i
		}
		right[i] = p
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := right[q[0]], left[q[1]]
		if l < r {
			ans[i] = pre[r] - pre[l]
		}
	}

	return ans
}
