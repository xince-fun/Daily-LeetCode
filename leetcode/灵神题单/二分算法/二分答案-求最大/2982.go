package main

func maximumLength(s string) int {
	n := len(s)
	cnt := make(map[byte][]int)
	for i, j := 0, 0; i < n; i = j {
		for j < n && s[j] == s[i] {
			j++
		}
		cnt[s[i]] = append(cnt[s[i]], j-i)
	}

	res := -1
	for _, vec := range cnt {
		left, right := 0, n
		for left+1 < right {
			mid := left + (right-left)/2
			count := 0
			for _, x := range vec {
				if x >= mid {
					count += x - mid + 1
				}
			}
			if count >= 3 {
				if mid > res {
					res = mid
				}
				left = mid
			} else {
				right = mid
			}
		}
	}
	return res
}
