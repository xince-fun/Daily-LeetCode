package main

func totalFruit(fruits []int) (ans int) {
	left := 0
	hash := make(map[int]int, 2)
	for right, x := range fruits {
		hash[x]++
		for len(hash) > 2 {
			hash[fruits[left]]--
			if hash[fruits[left]] == 0 {
				delete(hash, fruits[left])
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
