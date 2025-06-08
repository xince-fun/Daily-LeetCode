package main

// 最多有m个0和n个1 总长度不超过m+n
// 定义 f[i][j] 表示 容量为m n的背包

func findMaxForm(strs []string, m int, n int) int {
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for _, str := range strs {
		cnt1, cnt2 := 0, 0
		for _, ch := range str {
			if ch == '0' {
				cnt1++
			} else {
				cnt2++
			}
		}
		for i := m; i >= cnt1; i-- {
			for j := n; j >= cnt2; j-- {
				f[i][j] = max(f[i][j], f[i-cnt1][j-cnt2]+1)
			}
		}
	}
	return f[m][n]
}
