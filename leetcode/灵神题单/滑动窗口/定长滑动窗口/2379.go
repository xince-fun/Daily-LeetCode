package main

func minimumRecolors(blocks string, k int) (ans int) {
	ans = len(blocks)
	count := 0
	for i, in := range blocks {
		if in == 'B' {
			count++
		}
		if i < k-1 {
			continue
		}
		ans = min(ans, k-count)
		out := blocks[i-k+1]
		if out == 'B' {
			count--
		}
	}
	return
}
