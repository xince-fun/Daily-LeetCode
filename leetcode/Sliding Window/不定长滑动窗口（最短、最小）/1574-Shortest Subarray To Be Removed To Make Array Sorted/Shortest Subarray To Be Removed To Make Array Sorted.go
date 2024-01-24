package main

// 枚举左端点，移动右端点
func findLengthOfShortestSubarray1(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 { // 已经是非递减数组
		return 0
	}
	// 此时 arr[right-1] > arr[right]
	ans := right // 删除 arr[:right]
	for left := 0; left == 0 || arr[left-1] <= arr[left]; left++ {
		for right < n && arr[right] < arr[left] {
			right++
		}
		ans = min(ans, right-left-1) // 删除 arr[left+1:right]
	}
	return ans
}

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 { // 已经是非递减数组
		return 0
	}
	ans := right
	for left := 0; ; right++ {
		for right == n || arr[left] <= arr[right] {
			ans = min(ans, right-left+1)
			if arr[left] > arr[left+1] {
				return ans
			}
			left++
		}
	}
}
