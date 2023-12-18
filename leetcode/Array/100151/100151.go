package main

import (
	"sort"
)

/*
中位数贪心
给定一个数组a，那么取 a 的中位数x
x 到 a 中所有数的距离之和是最小的

1. 怎么预处理回文数？


2. 所有数改成哪个回文数最优？

*/

// func minimumCost(nums []int) int64 {
// 	// 对数组进行排序
// 	sort.Ints(nums)
// 	n := len(nums)

// 	// 找到中位数
// 	mid := nums[n/2]

// 	// 生成最接近中位数的小于等于它的回文数
// 	palindromeLess := generatePalindrome(mid, false)

// 	// 生成最接近中位数的大于等于它的回文数
// 	palindromeGreater := generatePalindrome(mid, true)

// 	// 计算总代价
// 	costLess := calculateCost(nums, palindromeLess)
// 	costGreater := calculateCost(nums, palindromeGreater)

// 	// 返回两种选择中的最小总代价
// 	if costLess < costGreater {
// 		return costLess
// 	}
// 	return costGreater
// }

// // 生成最接近中位数的回文数
// func generatePalindrome(mid int, greater bool) int {
// 	s := strconv.Itoa(mid)
// 	length := len(s)
// 	var palindrome string

// 	// 使用中位数的前半部分来构建回文数
// 	half := s[:length/2+length%2]
// 	palindrome = half + reverse(half[:length/2])

// 	// 将字符串回文数转换为整数
// 	res, _ := strconv.Atoi(palindrome)

// 	// 如果需要大于等于中位数的回文数，而生成的回文数小于中位数
// 	// 或者如果需要小于等于中位数的回文数，而生成的回文数大于中位数，则调整回文数
// 	if (greater && res < mid) || (!greater && res > mid) {
// 		if greater {
// 			// 大于等于时，递增前半部分
// 			newHalf, _ := strconv.Atoi(half)
// 			newHalf++
// 			half = strconv.Itoa(newHalf)
// 		} else {
// 			// 小于等于时，递减前半部分
// 			newHalf, _ := strconv.Atoi(half)
// 			newHalf--
// 			half = strconv.Itoa(newHalf)
// 		}
// 		// 重新构建回文数
// 		if length%2 == 0 {
// 			palindrome = half + reverse(half)
// 		} else {
// 			palindrome = half + reverse(half[:length/2])
// 		}
// 		res, _ = strconv.Atoi(palindrome)
// 	}

// 	return res
// }

// // 计算将所有元素变为指定值的总代价
// func calculateCost(nums []int, target int) int64 {
// 	var cost int64
// 	for _, num := range nums {
// 		cost += int64(abs(num - target))
// 	}
// 	return cost
// }

// // 求绝对值
// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// // 反转字符串
// func reverse(s string) string {
// 	r := []rune(s)
// 	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }

// 按顺序从小到大生成所有回文数
var pal = make([]int, 0, 109999)

func init() {
	// 按顺序从小到大生成所有回文数
	for base := 1; base <= 10000; base *= 10 {
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			pal = append(pal, x)
		}
		if base <= 1000 {
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				pal = append(pal, x)
			}
		}
	}
	pal = append(pal, 1_000_000_001) // 哨兵，防止下标越界

}

func minimumCost(nums []int) int64 {
	sort.Ints(nums)
	n := len(nums)

	cost := func(i int) (res int64) {
		target := pal[i]
		for _, x := range nums {
			res += int64(abs(x - target))
		}
		return
	}

	// 二分找中位数最近的回文数
	i := sort.SearchInts(pal, nums[(n-1)/2])
	if pal[i] <= nums[n/2] { // 回文数在中位数范围内
		return cost(i) // 直接变成pal[i]
	}

	return min(cost(i-1), cost(i))
}

// 求绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
