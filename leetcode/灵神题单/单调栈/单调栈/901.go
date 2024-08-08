package main

import "math"

/*
	type StockSpanner struct {
		cnt int
		arr []int
		st  []int
		sum []int
	}

	func Constructor() StockSpanner {
		return StockSpanner{
			cnt: 0,
			arr: make([]int, 0),
			st:  make([]int, 0),
			sum: make([]int, 0),
		}
	}

	func (t *StockSpanner) Next(price int) int {
		t.arr = append(t.arr, price)
		ans := 1
		for len(t.st) > 0 && price >= t.arr[t.st[len(t.st)-1]] {
			idx := t.st[len(t.st)-1]
			t.st = t.st[:len(t.st)-1]
			ans += t.sum[idx]
		}
		t.sum = append(t.sum, ans)
		t.st = append(t.st, t.cnt)
		t.cnt++
		return ans
	}
*/
type pair struct {
	day   int
	price int
}

type StockSpanner struct {
	st     []pair
	curDay int
}

func Constructor() StockSpanner {
	return StockSpanner{
		st:     []pair{{-1, math.MaxInt}},
		curDay: -1,
	}
}

func (s *StockSpanner) Next(price int) int {
	for price >= s.st[len(s.st)-1].price {
		s.st = s.st[:len(s.st)-1]
	}
	s.curDay++
	s.st = append(s.st, pair{s.curDay, price})
	return s.curDay - s.st[len(s.st)-2].day
}
