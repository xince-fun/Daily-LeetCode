package main

func countHomogenous(s string) (ans int) {
	for i, n := 0, len(s); i < n; {
		i0 := i
		for i++; i < n && s[i] == s[i-1]; i++ {
		}
		ans += (i - i0) * (i - i0 + 1) / 2 % 1_000_000_007
	}
	return
}
