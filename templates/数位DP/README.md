# 数位DP

```go
func countSpecialNumbers(n int) (ans int) {
	s := strconv.Itoa(n)
	m := len(s)
	memo := make([][1 << 10]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var f func(int, int, bool, bool) int
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1 // 得到了一个合法数字
			}
			return
		}
		if !isLimit && isNum {
			p := &memo[i][mask]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if mask>>d&1 == 0 { // d 不在 mask 中
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return f(0, 0, true, false)
}

```

模版 v2

```go
// 第一种写法（前导零不影响答案）
var f func(int, int, bool, bool) int
f = func(p, sum int, limitLow, limitHigh bool) (res int) {
	if p == n {
		// 不合法
		if sum > sumUpper {
			return 0
		}
		// 合法
		return 1
	}
	if !limitLow && !limitHigh {
		dv := &dp[p][sum]
		if *dv >= 0 {
			return *dv
		}
		defer func() { *dv = res }()
	}

	lo := 0
	if limitLow {
		lo = int(lowS[p] - '0')
	}
	// 注：不要修改这里！如果对数位有其它限制，应当写在下面 for 循环中
	hi := 9
	if limitHigh {
		hi = int(highS[p] - '0')
	}

	for d := lo; d <= hi; d++ {
		res += f(p+1, sum+d, limitLow && d == lo, limitHigh && d == hi)
		res %= mod
	}
	return
}
```