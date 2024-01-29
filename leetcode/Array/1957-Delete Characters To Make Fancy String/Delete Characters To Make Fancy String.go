package main

func makeFancyString(s string) string {
	arr := []byte{}
	n := len(s)
	i := 0
	for i < n {
		start := i
		for i < n && s[i] == s[start] {
			i++
		}
		arr = append(arr, s[start:min(i, start+2, n)]...)
	}
	return string(arr)
}
