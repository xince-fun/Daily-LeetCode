package main

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	count := 1
	index := 0
	for candies > 0 {
		give := 0
		if count >= candies {
			give = candies
		} else {
			give = count
		}
		ans[index%num_people] += give
		candies -= give
		count++
		index++
	}
	return ans
}
