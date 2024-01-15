package main

func buildNext(pattern string) []int {
	n := len(pattern)
	next := make([]int, n)
	next[0] = -1
	i, j := 0, -1
	for i < n-1 {
		if j == -1 || pattern[i] == pattern[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func kmpSearch(text, pattern string, next []int, k int) bool {
	m, n := len(pattern), len(text)
	i, j := 0, 0
	for i < n && j < m {
		if j == -1 || text[i] == pattern[j] {
			i++
			j++
		} else {
			j = next[j]
		}
		if j == m && i-j <= k {
			return true
		}
	}
	return false
}

func beautifulIndices(s string, a string, b string, k int) (ans []int) {
	nextA := buildNext(a)
	nextB := buildNext(b)
	n := len(s)
	for i := 0; i <= n-len(a); i++ {
		if s[i:i+len(a)] == a {
			for j := 0; j <= n-len(b); j++ {
				if s[j:j+len(b)] == b && abs(j-i) <= k {
					if kmpSearch(s, b, nextB, j) && kmpSearch(s, a, nextA, i) {
						ans = append(ans, i)
						break
					}
				}
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {

}
