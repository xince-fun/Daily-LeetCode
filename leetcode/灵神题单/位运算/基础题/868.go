package main

import "math/bits"

func binaryGap(n int) (ans int) {
	last := -1
	for i := 0; i < bits.Len(uint(n)); i++ {
		if n>>i&1 == 1 {
			if last != -1 {
				ans = max(ans, i-last)
			}
			last = i
		}
	}
	return
}
