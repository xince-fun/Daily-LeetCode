package main

// 类似于爬楼梯
// 2\3\4\5\6\8 一次爬 1-3阶
// 7\9 一次爬 1-4阶

// 最后爬了 x = mapping[nums[i]]阶， 就得先到 i - x阶
// dfs

const mod, mx int = 1e9 + 7, 1e5

var f = [mx + 1]int{1, 1, 2, 4}
var g = f

func init() {
	for i := 4; i <= mx; i++ { // 预处理所有长度的结果
		f[i] = (f[i-1] + f[i-2] + f[i-3]) % mod
		g[i] = (g[i-1] + g[i-2] + g[i-3] + g[i-4]) % mod
	}
}

func countTexts(s string) int {
	ans, cnt := 1, 0
	for i, c := range s {
		cnt++
		if i == len(s)-1 || s[i+1] != byte(c) { // 找到一个完整的组
			if c != '7' && c != '9' {
				ans = ans * f[cnt] % mod
			} else {
				ans = ans * g[cnt] % mod
			}
			cnt = 0
		}
	}
	return ans
}
