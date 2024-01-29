package main

func minCost(colors string, neededTime []int) (ans int) {
	for i, n := 0, len(colors); i < n; {
		i0 := i
		cnt := 0
		maxTime := 0
		for ; i < n && colors[i] == colors[i0]; i++ {
			maxTime = max(maxTime, neededTime[i])
			cnt += neededTime[i]
		}
		ans += cnt - maxTime
	}
	return
}
