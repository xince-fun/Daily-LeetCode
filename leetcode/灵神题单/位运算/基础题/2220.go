package main

func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func minBitFlips(start int, goal int) int {
	return onesCount(start ^ goal)
}
