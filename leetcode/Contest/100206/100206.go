package main

import (
	"fmt"
	"sort"
)

func maximumLength1(nums []int) (ans int) {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	ans = 1
	if count, ok := freq[1]; ok {
		if count%2 == 0 {
			count--
		}
		ans = max(ans, count)
	}
	keys := make([]int, 0)
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	dp := make(map[int]int)
	for _, num := range keys {
		numSquare := num * num
		if 2 <= freq[num] {
			if val, ok := dp[numSquare]; ok {
				dp[num] = 2 + val
				ans = max(ans, dp[num])
			} else {
				dp[num] = 1
			}
		} else {
			dp[num] = 1
		}
	}
	return ans
}

func maximumLength(nums []int) (ans int) {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	sort.Ints(nums)
	oneCnt := 0
	for _, num := range nums {
		cnt := 1
		if num == 1 {
			oneCnt++
			continue
		}
		for freq[num*num] > 0 {
			if freq[num] >= 2 {
				freq[num] -= 2
				cnt++
				num *= num
			} else {
				break
			}
		}
		ans = max(ans, cnt)
	}
	ans = ans*2 - 1
	if oneCnt%2 == 0 {
		oneCnt--
	}
	ans = max(ans, oneCnt)
	return
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	fmt.Println(maximumLength(nums))
}
