package main

func numOfSubarrays(arr []int, k int, threshold int) (ans int) {
	sum := 0
	for _, v := range arr[:k] {
		sum += v
	}
	if sum >= k*threshold {
		ans++
	}
	for i := k; i < len(arr); i++ {
		sum = sum - arr[i-k] + arr[i]
		if sum >= k*threshold {
			ans++
		}
	}
	return
}
