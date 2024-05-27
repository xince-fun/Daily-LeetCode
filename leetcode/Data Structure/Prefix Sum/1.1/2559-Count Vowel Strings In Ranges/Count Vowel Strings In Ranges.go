package main

func isVowelString(word string) bool {
	return isVowelLetter(word[0]) && isVowelLetter(word[len(word)-1])
}
func isVowelLetter(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func vowelStrings(words []string, queries [][]int) []int {
	s := make([]int, len(words)+1)
	for i, word := range words {
		s[i+1] = s[i]
		if isVowelString(word) {
			s[i+1]++
		}
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = s[query[1]+1] - s[query[0]]
	}
	return ans
}
