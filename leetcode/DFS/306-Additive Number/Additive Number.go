package main

import (
	"strconv"
)

// 逗号 选不选

func isAdditiveNumber1(num string) (ans bool) {
	n := len(num)

	path := make([]int, 0)
	var dfs func(int, int)
	dfs = func(i, start int) {
		if i == n {
			if len(path) >= 3 {
				ans = true
			}
			return
		}
		if i < n-1 {
			dfs(i+1, start)
		}
		// 选这个数
		s := num[start : i+1]
		str, _ := strconv.Atoi(s)
		if s != strconv.Itoa(str) {
			return
		}
		if len(path) >= 2 {
			if path[len(path)-1]+path[len(path)-2] != str {
				return
			}
		}
		path = append(path, str)
		dfs(i+1, i+1)
		path = path[:len(path)-1]
	}

	dfs(0, 0)
	return
}

// 从答案的角度来看

// 1. 当前操作? 构建一个下标 j >= i的数字
// 2. 子问题？ 从下标 >= i 中构造子集
// 3. 下一个子问题？ 从下表 >= j+1的数字中构造

func isAdditiveNumber(num string) (ans bool) {
	n := len(num)
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			if len(path) >= 3 {
				ans = true
			}
			return
		}
		for j := i; j < n; j++ {
			s := num[i : j+1]
			str, _ := strconv.Atoi(s)
			if s != strconv.Itoa(str) {
				break
			}
			if len(path) >= 2 {
				if path[len(path)-1]+path[len(path)-2] != str {
					continue
				}
			}
			path = append(path, str)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
