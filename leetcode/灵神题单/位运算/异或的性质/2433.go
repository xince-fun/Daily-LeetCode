package main

// pref[i] = arr[0] ^ arr[1] ^ ... ^ arr[i]

// pref[2] = (arr[0] ^ arr[1] ^ arr[2])
// arr[2] = (arr[0] ^ arr[1] ^ arr[2]) ^ (arr[0] ^ arr[1])

func findArray1(pref []int) []int {
	pre := pref[0]
	for i := 1; i < len(pref); i++ {
		pref[i] = pref[i] ^ pre
		pre ^= pref[i]
	}
	return pref
}

func findArray(pref []int) []int {
	for i := len(pref) - 1; i > 0; i-- {
		pref[i] = pref[i] ^ pref[i-1]
	}
	return pref
}
