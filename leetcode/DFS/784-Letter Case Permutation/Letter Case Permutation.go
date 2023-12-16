package main

import (
	"fmt"
)

// 当前操作？ 变不变这个元素
// 子问题？ 构造字符串>=i的部分
// 下一个子问题？ 构造字符串>=i+1的部分

func letterCasePermutation1(s string) (ans []string) {
	n := len(s)
	path := make([]byte, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		path[i] = s[i]
		dfs(i + 1)
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			path[i] = s[i] ^ 32
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}

// 从答案的角度来看

// 1. 当前操作？ 构造一个下标 j>=i的字符串
// 2. 子问题? 从下标 >= i 的字符串中构造子集
// 3. 下一个子问题？ 从下表 >= j+1的数字中构造子集

func letterCasePermutation(s string) (ans []string) {
	n := len(s)
	path := []byte(s)

	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, string(path))
		if i == n {
			return
		}
		for j := i; j < n; j++ {
			if (path[j] >= 'a' && path[j] <= 'z') || (path[j] >= 'A' && path[j] <= 'Z') {
				path[j] ^= 32 // 改变大小写
				dfs(j + 1)
				path[j] ^= 32 // 恢复
			}
		}
	}
	dfs(0)
	return
}

func main() {
	s := "a1b2"
	fmt.Println(letterCasePermutation(s))
}
