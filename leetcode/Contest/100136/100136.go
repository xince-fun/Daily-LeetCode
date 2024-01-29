package main

import "slices"

/*
 1. 相同数字一定要在同一个子数组中
 2. 推论：对于一个数字，找到它最左边的下标 p 和最右边的下标 q
    那么 [p,q] 是不能分割的
 3. 56 - 区间合并

算法
1. 遍历数组，维护每个元素最后一次出现的下标
2. 再次遍历数组，去合并区间
*/
const mod int = 1e9 + 7

func numberOfGoodPartitions1(nums []int) int {
	r := map[int]int{}
	for i, num := range nums {
		r[num] = i
	}
	m, maxr := 0, 0
	for i, num := range nums {
		maxr = max(maxr, r[num])
		if maxr == i {
			m += 1
		}
	}
	return pow(2, m-1)
}

func numberOfGoodPartitions(nums []int) int {
	type pair struct{ l, r int }
	ps := map[int]pair{}
	for i, x := range nums {
		if p, ok := ps[x]; ok {
			p.r = i
			ps[x] = p
		} else {
			ps[x] = pair{i, i}
		}
	}

	a := make([]pair, 0, len(ps))
	for _, p := range ps {
		a = append(a, p)
	}
	slices.SortFunc(a, func(a, b pair) int { return a.l - b.l }) // 按区间左端点排序

	ans := 1
	maxR := a[0].r
	for _, p := range a[1:] {
		if p.l > maxR {
			ans = ans * 2 % 1_000_000_007
		}
		maxR = max(maxR, p.r)
	}
	return ans
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
