package main

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	// 假定 len(nums1) < len(nums2)
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays1(nums2, nums1)
	}
	left, right, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for left <= right {
		//
		nums1Mid = left + (right-left)>>1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			right = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			left = nums1Mid + 1
		} else {
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft + midRight)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 0 {
		left := find(nums1, nums2, 0, 0, n/2)
		right := find(nums1, nums2, 0, 0, n/2+1)
		return (float64(left) + float64(right)) / 2
	} else {
		return float64(find(nums1, nums2, 0, 0, n/2+1))
	}
}

// 找到两个数组中第k大的元素
func find(nums1, nums2 []int, i, j, k int) int {
	if len(nums1)-i > len(nums2)-j {
		return find(nums2, nums1, j, i, k)
	}
	if len(nums1) == i {
		return nums2[j+k-1]
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	idx1, idx2 := min(len(nums1), i+k/2), j+k-k/2
	if nums1[idx1-1] < nums2[idx2-1] {
		return find(nums1, nums2, idx1, j, k-(idx1-i))
	} else {
		return find(nums1, nums2, i, idx2, k-(idx2-j))
	}
}

func main() {

}
