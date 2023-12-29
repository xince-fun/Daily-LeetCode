package sum

import (
	"sort"
)

func fourSum(nums []int, target int) (ans [][]int) {
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		x := nums[i]
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if x+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			y := nums[j]
			if j > i+1 && y == nums[j-1] {
				continue
			}
			if x+y+nums[j+1]+nums[j+2] > target {
				break
			}
			if x+y+nums[n-2]+nums[n-1] < target {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					right--
					for right > left && nums[right] == nums[right+1] {
						right--
					}
				}
			}
		}
	}
	return
}
