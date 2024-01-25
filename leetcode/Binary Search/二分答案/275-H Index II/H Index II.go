package main

// 闭区间写法
func hIndex1(citations []int) int {
	// 在区间 [left, right] 内询问
	left, right := 1, len(citations)
	for left <= right {
		// 循环不变量：
		// left-1 的回答一定是「是」
		// right+1 的回答一定为「否」
		mid := (left + right) / 2
		// 引用次数最多的 mid 篇文章， 引用次数均 >= mid
		if citations[len(citations)-mid] >= mid {
			left = mid + 1 // 询问范围缩小到 [mid+1, right]
		} else {
			right = mid - 1 // 询问范围缩小到 [left, mid-1]
		}
	}
	// 循环结束后 right 等于 left-1，回答一定为「是」
	// 根据循环不变量，right 现在是最大的回答为「是」的数
	return right
}

// 左闭右开区间写法
func hIndex2(citations []int) int {
	// 在区间 [left, right) 内询问
	left, right := 1, len(citations)+1
	for left < right { // 区间不为空
		// 循环不变量：
		// left-1 的回答一定是「是」
		// right 的回答一定是「否」
		mid := (left + right) / 2
		// 引用次数最多的 mid 篇文章， 引用次数均 >= mid
		if citations[len(citations)-mid] >= mid {
			left = mid + 1 // 询问范围缩小到 [mid+1, right)
		} else {
			right = mid // 询问范围缩小到 [left, mid)
		}
	}
	// 循环结束后 left 等于 right，回答一定是「否」
	// 根据循环不变量，left-1 现在是最大的回答为「是」的数
	return left - 1
}

// 左开右闭区间写法
func hIndex3(citations []int) int {
	// 在区间 (left, right] 内询问
	left, right := 0, len(citations)
	for left < right { // 区间不为空
		// 循环不变量：
		// left 的回答一定是「是」
		// right+1 的回答一定是「否」
		mid := (left + right + 1) / 2 // 保证 mid 在二分区间内
		// 引用次数最多的 mid 篇文章， 引用次数均 >= mid
		if citations[len(citations)-mid] >= mid {
			left = mid // 询问范围缩小到 (mid, right]
		} else {
			right = mid - 1 // 询问范围缩小到 (left, mid-1]
		}
	}
	// 循环结束后 left 等于 right，回答一定是「是」
	// 根据循环不变量，left 现在是最大的回答为「是」的数
	return left
}

// 开区间写法
func hIndex(citations []int) int {
	// 在区间 (left, right) 内询问
	left, right := 0, len(citations)+1
	for left+1 < right { // 区间不为空
		// 循环不变量：
		// left 的回答一定是「是」
		// right 的回答一定是「否」
		mid := (left + right + 1) / 2 // 保证 mid 在二分区间内
		// 引用次数最多的 mid 篇文章， 引用次数均 >= mid
		if citations[len(citations)-mid] >= mid {
			left = mid // 询问范围缩小到 (mid, right)
		} else {
			right = mid // 询问范围缩小到 (left, mid)
		}
	}
	// 循环结束后 left+1 等于 right，回答一定是「否」
	// 根据循环不变量，left 现在是最大的回答为「是」的数
	return left
}
