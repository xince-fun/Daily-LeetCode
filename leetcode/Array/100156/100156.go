package main

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	if source == target {
		return 0
	}

	n, m := len(source), len(target)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}

}
