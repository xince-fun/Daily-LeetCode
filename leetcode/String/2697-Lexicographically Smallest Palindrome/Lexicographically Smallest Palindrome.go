package main

import "fmt"

func makeSmallestPalindrome(s string) string {
	arr := []byte(s)
	for l, r := 0, len(arr)-1; l <= r; l, r = l+1, r-1 {
		if arr[l] > arr[r] {
			arr[l] = arr[r]
		} else {
			arr[r] = arr[l]
		}
	}
	return string(arr)
}

func main() {
	str := "egcefe"
	fmt.Println(makeSmallestPalindrome(str))
}
