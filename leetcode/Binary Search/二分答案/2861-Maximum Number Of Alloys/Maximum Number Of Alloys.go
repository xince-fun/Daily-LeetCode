package main

import (
	"fmt"
	"sort"
)

func maxNumberOfAlloys1(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	left, right := 0, int(1e9)

	check := func(mid int) bool {
		for i := 0; i < k; i++ {
			tot := 0
			for j := 0; j < n && tot <= budget; j++ {
				if stock[j] <= composition[i][j]*mid {
					tot += (composition[i][j]*mid - stock[j]) * cost[j]
				}
			}
			if tot <= budget {
				return true
			}
		}
		return false
	}

	for left <= right {
		// 循环不变量：
		// left-1 一定是「是」
		// right+1 一定是「否」
		// mid 为制造的合金数
		mid := left + (right-left)/2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func maxNumberOfAlloys2(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	return sort.Search(1e9, func(mid int) bool {
		for i := 0; i < k; i++ {
			tot := 0
			for j := 0; j < n && tot <= budget; j++ {
				if stock[j] <= composition[i][j]*mid {
					tot += (composition[i][j]*mid - stock[j]) * cost[j]
				}
			}
			if tot <= budget {
				return false
			}
		}
		return true
	}) - 1
}

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) (ans int) {
	mx := stock[0]
	for _, s := range stock {
		mx = min(mx, s)
	}
	mx += budget
	for _, com := range composition {
		res := sort.Search(mx, func(num int) bool {
			num++
			money := 0
			for i, s := range stock {
				if s < com[i]*num {
					money += (com[i]*num - s) * cost[i]
					if money > budget {
						return true
					}
				}
			}
			return false
		})
		ans = max(ans, res)
	}
	return
}

func main() {
	n := 3
	k := 2
	budget := 15
	composition := [][]int{{1, 1, 1}, {1, 1, 10}}
	stock := []int{0, 0, 0}
	cost := []int{1, 2, 3}
	fmt.Println(maxNumberOfAlloys(n, k, budget, composition, stock, cost))
}
