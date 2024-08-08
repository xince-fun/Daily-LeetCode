package main

func minimumSteps(s string) (ans int64) {
	// 100100
	// 1 + 1 + 2 + 2 = 6
	// 101
	// 1
	// 100101
	// 1 + 1 + 2 = 4
	// 100
	// 1 + 1  = 2
	// 0111
	// 0
	count := 0
	for _, c := range s {
		if c == '1' {
			count++
		} else {
			ans += int64(count)
		}
	}
	return ans
}
