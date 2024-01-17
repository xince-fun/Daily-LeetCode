package main

import (
	"fmt"
	"strconv"
)

func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	memo := make([][1 << 10]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	// 返回从 i 开始填数字，i 前面填的数字的集合是 mask，能构造出的特殊整数的树木
	// isLimit 表示前面填的数字是否都是 n 对应位上的， 如果为true， 那么当前位至多为 int(s[i])，否则至多为 '9'
	// isNum 表示前面是否填了数字（是否跳过），如果为 true， 那么当前位可以从 0开始，如果为 false，那么我们可以跳过，或者从 1开始填数字
	var f func(int, int, bool, bool) int
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}
		if !isLimit && isNum {
			p := &memo[i][mask]
			if *p >= 0 {
				fmt.Printf("%b\n", mask)
				fmt.Println(i, *p)
				return *p
			}
			defer func() { *p = res }()
		}
		if !isNum {
			res += f(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for ; d <= up; d++ {
			if mask>>d&1 == 0 {
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return f(0, 0, true, false)
}

func main() {
	fmt.Println(countSpecialNumbers(420))
}
