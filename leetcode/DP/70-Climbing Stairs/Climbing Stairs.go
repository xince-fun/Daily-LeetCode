package main

// f[i] = f[i-1], f[i-2]
// f[i+2] = f[i+1], f[i]

// 当前=f0+f1
// f0表示上上一个，f1表示上一个
// newF=f0+f1
// f0=f1
// f1=newF

func climbStairs1(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func climbStairs2(n int) int {
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		newF := f0 + f1
		f0 = f1
		f1 = newF
	}
	return f1
}

func climbStairs(n int) int {
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		f0, f1 = f1, f1+f0
	}
	return f1
}
