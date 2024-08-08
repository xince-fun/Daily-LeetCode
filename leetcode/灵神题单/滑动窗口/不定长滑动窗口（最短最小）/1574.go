package main

func findLengthOfShortestSubarray(arr []int) (ans int) {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}
	if right == 0 { // 此时已经是非递减
		return 0
	}
	// 要删除 arr[:right]
	ans = right
	for left := 0; left == 0 || arr[left-1] <= arr[left]; left++ {
		for right < n && arr[right] < arr[left] {
			right++
		}
		ans = min(ans, right-left-1)
	}
	return
}
