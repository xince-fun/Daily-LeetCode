package main

func sumOddLengthSubarrays(arr []int) int {
	s := make([]int, len(arr)+1)
	for i, x := range arr {
		s[i+1] = s[i] + x
	}
	ans := 0

	for length := 1; length <= len(arr); length += 2 {
		for l := 0; l+length-1 < len(arr); l++ {
			r := l + length
			ans += s[r] - s[l]
		}
	}
	return ans
}
