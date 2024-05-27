package main

func missingRolls(rolls []int, mean int, n int) []int {
	tot := (len(rolls) + n) * mean
	sum := 0
	for _, x := range rolls {
		sum += x
	}
	ans := make([]int, n)
	missing := tot - sum

	avg := missing / n
	if avg < 1 || (avg > 6) || (avg == 6 && missing%n != 0) {
		return []int{}
	}
	remain := missing - avg*n
	for i := 0; i < n; i++ {
		ans[i] = avg
		if ans[i] < 6 && remain > 0 {
			ans[i] += 1
			remain--
		}
	}
	return ans
}
