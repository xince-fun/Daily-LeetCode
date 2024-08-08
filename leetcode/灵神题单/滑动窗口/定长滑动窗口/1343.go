package main

func numOfSubarrays(arr []int, k int, threshold int) (ans int) {
	sum := 0
	for i := 0; i < k-1; i++ {
		sum += arr[i]
	}
	for i := k - 1; i < len(arr); i++ {
		sum += arr[i]
		if float64(sum)/float64(k) > float64(threshold) {
			ans++
		}
		sum -= arr[i-k+1]
	}
	return
}
