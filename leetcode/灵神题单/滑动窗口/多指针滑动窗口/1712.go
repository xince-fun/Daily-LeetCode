package main

func waysToSplit(nums []int) (ans int) {
	const mod = 1e9 + 7
	n := len(nums)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	// sum(left) = pre[l]
	// sum(mid) = pre[r]-pre[l]
	// sum(right) = pre[n]-pre[r]

	// pre[l] <= pre[r] - pre[l]
	// pre[r]-pre[l] <= pre[n]-pre[r]

	left1, left2 := 1, 1
	for right := 2; right < n; right++ {
		for left1 < right && pre[left1] <= pre[right]-pre[left1] {
			left1++
		}
		for left2 < left1 && pre[right]-pre[left2] > pre[n]-pre[right] {
			left2++
		}
		ans += left1 - left2
	}
	return
}
