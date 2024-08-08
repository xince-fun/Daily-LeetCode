package main

func xorQueries(arr []int, queries [][]int) []int {
	pre := make([]int, len(arr)+1)
	ans := make([]int, len(queries))
	for i, x := range arr {
		pre[i+1] = pre[i] ^ x
	}

	for i, q := range queries {
		ans[i] = pre[q[1]+1] ^ pre[q[0]]
	}
	return ans
}
