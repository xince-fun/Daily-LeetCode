package main

func totalFruit(fruits []int) (ans int) {
	n := len(fruits)
	sum, count := 0, 0
	left := 0
	cnt := make([]int, n)
	for _, fruit := range fruits {
		sum += 1
		if cnt[fruit] == 0 {
			count++
		}
		cnt[fruit]++
		for count > 2 {
			cnt[fruits[left]]--
			if cnt[fruits[left]] == 0 {
				count--
			}
			left++
			sum--
		}
		ans = max(ans, sum)
	}
	return
}
