package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func countPowerfulIntegers(level, isStartInclusive, isFinishInclusive int, start,
// 	finish, S string, limit, length int, dp [][][]int64) int64 {
// 	if level == len(finish) {
// 		return 1
// 	}
// 	if dp[level][isStartInclusive][isFinishInclusive] != -1 {
// 		return dp[level][isStartInclusive][isFinishInclusive]
// 	}
// 	var count int64 = 0
// 	startDigit := 0
// 	if isStartInclusive != 0 {
// 		startDigit = int(start[level] - '0')
// 	}
// 	finishDigit := limit
// 	if isFinishInclusive != 0 {
// 		finishDigit = min(int(finish[level]-'0'), limit)
// 	}

// 	for i := startDigit; i <= finishDigit; i++ {
// 		newStartInclusive := 0
// 		if i == int(start[level]-'0') {
// 			newStartInclusive = isStartInclusive
// 		}
// 		newFinishInclusive := 0
// 		if i == int(finish[level]-'0') {
// 			newFinishInclusive = isFinishInclusive
// 		}

// 		if level < length {
// 			count += countPowerfulIntegers(level+1, newStartInclusive, newFinishInclusive, start, finish, S, limit, length, dp)
// 		} else {
// 			if int(S[level-length]-'0') == i {
// 				count += countPowerfulIntegers(level+1, newStartInclusive, newFinishInclusive, start, finish, S, limit, length, dp)
// 			}
// 		}
// 	}
// 	dp[level][isStartInclusive][isFinishInclusive] = count
// 	return count
// }

// func numberOfPowerfulInt(start, finish int64, limit int, s string) int64 {
// 	startStr := strconv.FormatInt(start, 10)
// 	finishStr := strconv.FormatInt(finish, 10)
// 	startStr = strings.Repeat("0", max(0, len(finishStr)-len(startStr))) + startStr
// 	dp := make([][][]int64, 16)
// 	for i := range dp {
// 		dp[i] = make([][]int64, 2)
// 		for j := range dp[i] {
// 			dp[i][j] = make([]int64, 2)
// 			for k := range dp[i][j] {
// 				dp[i][j][k] = -1
// 			}
// 		}
// 	}
// 	lim := limit
// 	length := len(finishStr) - len(s)
// 	return countPowerfulIntegers(0, 1, 1, startStr, finishStr, s, lim, length, dp)
// }

/*
数位DP通用模板 v2.0

[low, high]



*/

func numberOfPowerfulInt1(start, finish int64, limit int, s string) int64 {
	low := strconv.FormatInt(start, 10)
	high := strconv.FormatInt(finish, 10)
	n := len(high)

	low = strings.Repeat("0", n-len(low)) + low // 补前导零，和 high 对齐
	diff := n - len(s)

	memo := make([]int64, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int, bool, bool) int64
	dfs = func(i int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			return 1
		}

		if !limitLow && !limitHigh {
			p := &memo[i]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow {
			lo = int(low[i] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(high[i] - '0')
		}

		if i < diff {
			for d := lo; d <= min(hi, limit); d++ {
				res += dfs(i+1, limitLow && d == lo, limitHigh && d == hi)
			}
		} else {
			x := int(s[i-diff] - '0')
			if lo <= x && x <= min(hi, limit) {
				res += dfs(i+1, limitLow && x == lo, limitHigh && x == hi)
			}
		}
		return
	}
	return dfs(0, true, true)
}

func numberOfPowerfulInt(start, finish int64, limit int, s string) int64 {
	low := strconv.FormatInt(start, 10)
	high := strconv.FormatInt(finish, 10)
	n := len(high)

	low = strings.Repeat("0", n-len(low)) + low // 补前导零，和 high 对齐
	diff := n - len(s)

	memo := make([]int64, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int, bool, bool, bool) int64
	dfs = func(i int, limitLow, limitHigh, isNum bool) (res int64) {
		if i == n {
			if isNum {
				return 1
			}
			return 0
		}

		if !isNum && low[i] == '0' {
			res += dfs(i+1, true, false, false)
		}

		if !limitLow && !limitHigh {
			p := &memo[i]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow {
			lo = int(low[i] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(high[i] - '0')
		}

		d0 := 1
		if isNum {
			d0 = 0
		}
		if i < diff {
			for d := max(lo, d0); d <= min(hi, limit); d++ {
				res += dfs(i+1, limitLow && d == lo, limitHigh && d == hi, true)
			}
		} else {
			x := int(s[i-diff] - '0')
			if max(lo, d0) <= x && x <= min(hi, limit) {
				res += dfs(i+1, limitLow && x == lo, limitHigh && x == hi, true)
			}
		}
		return
	}
	return dfs(0, true, true, false)
}

func main() {
	fmt.Println(numberOfPowerfulInt(20, 1159, 5, "20"))
}
