package main

const mod = 1e9 + 7

func numOfSubarrays(arr []int) (ans int) {
	n := len(arr)
	sum := 0
	m := make(map[int]int)
	m[0%2] = 1

	for i := 0; i < n; i++ {
		sum += arr[i]
		ans += m[(sum+1)%2]
		m[sum%2]++
	}
	return ans % mod
}
