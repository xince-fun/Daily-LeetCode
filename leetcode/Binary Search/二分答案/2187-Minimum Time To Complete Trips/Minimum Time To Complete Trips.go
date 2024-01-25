package main

import (
	"fmt"
	"sort"
)

/*
选择时间1的时候，完成的趟数为 1/1 + 1/2 + 1/3 = 1
当选择时间2的时候，完成的趟数为 2/1 + 2/2 + floor(2/3) = 3
时间3的时候，完成的趟数为 3/1 + floor(3/2) + floor(3/3) = 3+1+1=5
时间4的时候，完成的趟数为 4/1 + 4/2 + 4/3 = 4+2+1=7
上届为无穷
*/

// 闭区间
func minimumTime1(time []int, totalTrips int) int64 {
	// 区间为 [left, right]
	left := 1
	maxTime := time[0]
	for _, t := range time {
		if t > maxTime {
			maxTime = t
		}
	}
	right := maxTime * totalTrips
	cal := func(div int) (sum int) {
		for _, t := range time {
			sum += div / t
		}
		return
	}
	for left <= right {
		// 循环不变量
		// left-1的回答一定是「否」
		// right+1的回答一定是「是」
		mid := left + (right-left)/2
		// mid时间，如果趟数小于totalTrips
		if cal(mid) < totalTrips {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// right 等于 left-1
	// 根据循环不变量， right 现在是最大的回答为「是」的数
	return int64(right) + 1
}

// 左闭右开区间
func minimumTime2(time []int, totalTrips int) int64 {
	// 区间为 [left, right)
	left := 1
	maxTime := time[0]
	for _, t := range time {
		if t > maxTime {
			maxTime = t
		}
	}
	right := maxTime*totalTrips + 1
	cal := func(div int) (sum int) {
		for _, t := range time {
			sum += div / t
		}
		return
	}
	for left < right {
		// 循环不变量
		// left-1的回答一定是「否」
		// right的回答一定是「是」
		mid := left + (right-left)/2
		// mid时间，如果趟数小于totalTrips
		if cal(mid) < totalTrips {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// right 等于 left
	// 根据循环不变量， right 现在是最大的回答为「是」的数
	return int64(right)
}

// 左开右闭区间
func minimumTime3(time []int, totalTrips int) int64 {
	// 区间为 (left, right]
	left := 0
	maxTime := time[0]
	for _, t := range time {
		if t > maxTime {
			maxTime = t
		}
	}
	right := maxTime * totalTrips
	cal := func(div int) (sum int) {
		for _, t := range time {
			sum += div / t
		}
		return
	}
	for left < right {
		// 循环不变量
		// left的回答一定是「否」
		// right+1的回答一定是「是」
		mid := left + (right-left)/2
		// mid时间，如果趟数小于totalTrips
		if cal(mid) < totalTrips {
			left = mid
		} else {
			right = mid - 1
		}
	}
	// right 等于 left
	// 根据循环不变量， right 现在是最大的回答为「是」的数
	return int64(right) + 1
}

// 开区间
func minimumTime4(time []int, totalTrips int) int64 {
	// 区间为 (left, right)
	left := 0
	maxTime := time[0]
	for _, t := range time {
		if t > maxTime {
			maxTime = t
		}
	}
	right := maxTime*totalTrips + 1
	cal := func(div int) (sum int) {
		for _, t := range time {
			sum += div / t
		}
		return
	}
	for left+1 < right {
		// 循环不变量
		// left的回答一定是「否」
		// right的回答一定是「是」
		mid := left + (right-left+1)/2
		// mid时间，如果趟数小于totalTrips
		if cal(mid) < totalTrips {
			left = mid
		} else {
			right = mid
		}
	}
	// right 等于 left+1
	// 根据循环不变量， right 现在是最大的回答为「是」的数
	return int64(right)
}

// 用sort.Search
func minimumTime5(time []int, totalTrips int) int64 {
	return int64(sort.Search(totalTrips*1e7, func(i int) bool {
		tot := 0
		for _, t := range time {
			tot += i / t
			if tot >= totalTrips {
				return true
			}
		}
		return false
	}))
}

func minimumTime(time []int, totalTrips int) int64 {
	maxTime := time[0]
	for _, t := range time {
		if t > maxTime {
			maxTime = t
		}
	}
	return int64(sort.Search(totalTrips*maxTime, func(i int) bool {
		tot := 0
		for _, t := range time {
			tot += i / t
			if tot >= totalTrips {
				return true
			}
		}
		return false
	}))
}

func main() {
	time := []int{1, 2, 3}
	totalTrips := 5

	fmt.Println(minimumTime(time, totalTrips))
}
