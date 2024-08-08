package main

func maxScore1(cardPoints []int, k int) (ans int) {
	total := 0
	for _, num := range cardPoints {
		total += num
	}
	sum := 0
	ans = total
	m := len(cardPoints) - k
	if m == 0 {
		return total
	}
	for i, in := range cardPoints {
		sum += in
		if i < m-1 {
			continue
		}
		ans = min(ans, sum)
		out := cardPoints[i-m+1]
		sum -= out
	}
	ans = total - ans
	return
}

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	m := n - k
	s := 0
	for _, num := range cardPoints[:m] {
		s += num
	}
	total := s
	minS := s
	for i := m; i < n; i++ {
		total += cardPoints[i]
		s += cardPoints[i] - cardPoints[i-m]
		minS = min(minS, s)
	}
	return total - minS
}
