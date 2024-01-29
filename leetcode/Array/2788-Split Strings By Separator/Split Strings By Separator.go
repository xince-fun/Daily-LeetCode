package main

func splitWordsBySeparator(words []string, separator byte) (ans []string) {
	for _, w := range words {
		// 分组循环
		for i, n := 0, len(w); i < n; {
			if w[i] == separator {
				i++
				continue
			}
			start := i
			for i++; i < n && w[i] != separator; i++ {
			}
			ans = append(ans, w[start:i])
		}
	}
	return
}
