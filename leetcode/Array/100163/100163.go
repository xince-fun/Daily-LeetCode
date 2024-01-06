package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countPowerfulIntegers(level, isStartInclusive, isFinishInclusive int, start,
	finish, S string, limit, length int, dp [][][]int64) int64 {
	if level == len(finish) {
		return 1
	}
	if dp[level][isStartInclusive][isFinishInclusive] != -1 {
		return dp[level][isStartInclusive][isFinishInclusive]
	}
	var count int64 = 0
	startDigit := 0
	if isStartInclusive != 0 {
		startDigit = int(start[level] - '0')
	}
	finishDigit := limit
	if isFinishInclusive != 0 {
		finishDigit = min(int(finish[level]-'0'), limit)
	}

	for i := startDigit; i <= finishDigit; i++ {
		newStartInclusive := 0
		if i == int(start[level]-'0') {
			newStartInclusive = isStartInclusive
		}
		newFinishInclusive := 0
		if i == int(finish[level]-'0') {
			newFinishInclusive = isFinishInclusive
		}

		if level < length {
			count += countPowerfulIntegers(level+1, newStartInclusive, newFinishInclusive, start, finish, S, limit, length, dp)
		} else {
			if int(S[level-length]-'0') == i {
				count += countPowerfulIntegers(level+1, newStartInclusive, newFinishInclusive, start, finish, S, limit, length, dp)
			}
		}
	}
	dp[level][isStartInclusive][isFinishInclusive] = count
	return count
}

func numberOfPowerfulInt(start, finish int64, limit int, s string) int64 {
	startStr := strconv.FormatInt(start, 10)
	finishStr := strconv.FormatInt(finish, 10)
	startStr = strings.Repeat("0", max(0, len(finishStr)-len(startStr))) + startStr
	dp := make([][][]int64, 16)
	for i := range dp {
		dp[i] = make([][]int64, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int64, 2)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	lim := limit
	length := len(finishStr) - len(s)
	return countPowerfulIntegers(0, 1, 1, startStr, finishStr, s, lim, length, dp)
}

func main() {
	fmt.Println(numberOfPowerfulInt(20, 1159, 5, "20"))
}
