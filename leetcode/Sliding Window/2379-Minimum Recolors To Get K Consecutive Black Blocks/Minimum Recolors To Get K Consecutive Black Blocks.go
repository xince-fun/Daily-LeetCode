package main

import "strings"

func minimumRecolors(blocks string, k int) int {
	cntW := strings.Count(blocks[:k], "W")
	ans := cntW
	for i := k; i < len(blocks); i++ {
		cntW += int(blocks[i]&1) - int(blocks[i-k]&1)
		ans = min(ans, cntW)
	}
	return ans
}
