package main

import "slices"

// 定义 f[i] 表示为 s[i] 结尾的最大价值， 不和左边拼起来就是 vals[i] 或者 下标[i], 和 i左边拼起来就是 f[i-1] + (vals[i] or 下标i)

func maximumCostSubstring1(s string, chars string, vals []int) int {
	n := len(s)
	m := map[byte]int{}
	for i, x := range vals {
		m[chars[i]] = x
	}
	f := make([]int, n)
	if v, ok := m[s[0]]; ok {
		f[0] = max(v, 0)
	} else {
		f[0] = int(s[0]-'a') + 1
	}
	for i := 1; i < n; i++ {
		tmp := 0
		if v, ok := m[s[i]]; ok {
			tmp = v
		} else {
			tmp = int(s[i]-'a') + 1
		}
		f[i] = max(f[i-1], 0) + tmp
	}
	return slices.Max(f)
}

func maximumCostSubstring2(s string, chars string, vals []int) (ans int) {
	m := map[byte]int{}
	for i, x := range vals {
		m[chars[i]] = x
	}
	f := 0
	for i := 0; i < len(s); i++ {
		tmp := 0
		if v, ok := m[s[i]]; ok {
			tmp = v
		} else {
			tmp = int(s[i]-'a') + 1
		}
		f = max(f, 0) + tmp
		ans = max(ans, f)
	}
	return
}

func maximumCostSubstring(s string, chars string, vals []int) (ans int) {
	mapping := [26]int{}
	for i := range mapping {
		mapping[i] = i + 1
	}
	for i, c := range chars {
		mapping[c-'a'] = vals[i]
	}
	f := 0
	for _, c := range s {
		f = max(f, 0) + mapping[c-'a']
		ans = max(ans, f)
	}
	return
}
