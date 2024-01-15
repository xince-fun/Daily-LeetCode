package main

import (
	"strings"
)

// findAll returns all the occurrences of a substring in a given string
func findAll(s string, sub string) []int {
	start := 0
	var result []int

	for {
		index := strings.Index(s[start:], sub)
		if index == -1 {
			break
		}
		result = append(result, start+index)
		start += index + len(sub)
	}

	return result
}

// beautifulIndices finds all beautiful indices in the string s
func beautifulIndices(s string, a string, b string, k int) []int {
	iCandidates := findAll(s, a)
	jCandidates := findAll(s, b)

	var results []int

	for _, i := range iCandidates {
		for _, j := range jCandidates {
			if abs(j-i) <= k {
				results = append(results, i)
				break
			}
		}
	}

	return results
}

// beautifulIndicesOld is another variant of the beautiful indices function
func beautifulIndicesOld(s string, a string, b string, k int) []int {
	i := 0
	j := 0
	var iCandidates []int
	var jCandidates []int

	for i < len(s) || j < len(s) {
		if strings.HasPrefix(s[i:], a) && i <= len(s)-len(a) {
			iCandidates = append(iCandidates, i)
			i += len(a)
		} else {
			i++
		}

		if strings.HasPrefix(s[j:], b) && j <= len(s)-len(b) {
			jCandidates = append(jCandidates, j)
			j += len(b)
		} else {
			j++
		}
	}

	var results []int

	for _, i := range iCandidates {
		for _, j := range jCandidates {
			if abs(j-i) <= k {
				results = append(results, i)
				break
			}
		}
	}

	return results
}

// abs returns the absolute difference between two integers
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// main function for testing
func main() {
	s := "isawsquirrelnearmysquirrelhouseohmy"
	a := "my"
	b := "squirrel"
	k := 15

	indices := beautifulIndices(s, a, b, k)
	// Use indices as needed, e.g., print them
	for _, index := range indices {
		println(index)
	}

	// You can also test the beautifulIndicesOld function similarly
}
