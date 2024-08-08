package main

import (
	"slices"
	"sort"
)

// sort.Search 的使用技巧·其一
// sort.Search(n, f) 需要满足当 x 从小到大时，f(x) 先 false 后 true
// 若 f(x) 是先 true 后 false，且目标是找到最大的使 f(x) 为 true 的 x
// 这种情况可以考虑二分 !f(x)，则二分结果是最小的使 f(x) 为 false 的 x，将其 -1 就得到了最大的使 f(x) 为 true 的 x
// 由于要对结果 -1，sort.Search 传入的上界需要 +1
// 更加简单的写法是，在 f(x) 内部将 x++，这样就不需要对上界和结果调整 ±1 了
//
// 下面以二分求 int(sqrt(90)) 为例来说明这一技巧
// 这相当于求最大的满足 x*x<=90 的 x
// 于是定义 f(x) 返回 x*x<=90，注意这是一个先 true 后 false 的 f(x)
// 我们可以改为判断 f(x+1)，即用 f(x+1) 的返回结果代替 f(x) 的返回结果
// 同时，将 f(x) 改为先 false 后 true，即返回 x*x>90
// 这样二分的结果就恰好停在最大的满足原 f(x) 为 true 的 x 上

// sort.Search(10, func(x int) bool {
// 	x++
// 	return x*x > 90
// })

// sort.Search 的使用技巧·其二
// 若要求出一个和二分结果相关的东西
// 可以在返回值为 true 时记录下相关数据（若有多个地方返回 true，可以用 defer 来简化）
// 这样可以避免在二分结束后再计算一次
// 为了保证能至少触发一次 true，某些情况下需要将二分上界 +1
// https://codeforces.com/problemset/problem/1100/E

func maximumCandies1(candies []int, k int64) int {
	left, right := 0, slices.Max(candies)+1
	for left+1 < right {
		mid := left + (right-left)/2
		cnt := 0
		for _, c := range candies {
			cnt += c / mid
		}
		if int64(cnt) >= k {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func maximumCandies(candies []int, k int64) int {
	return sort.Search(slices.Max(candies)+1, func(i int) bool {
		i++
		cnt := int64(0)
		for _, c := range candies {
			cnt += int64(c / i)
		}
		return cnt < k
	})
}
