package main

import "slices"

func maximumLength1(s string) (ans int) {
	m := map[string]int{}
	n := len(s)
	left, right := 0, 0
	for left < n {
		sub := []byte{}
		sub = append(sub, s[left])
		m[string(sub)]++
		if m[string(sub)] >= 3 {
			ans = max(ans, 1)
		}
		right = left + 1
		for right < n && s[right] == s[right-1] {
			sub = append(sub, s[right])
			right++
			m[string(sub)]++
			if m[string(sub)] >= 3 {
				ans = max(ans, right-left)
			}
		}
		left++
	}
	return
}

// 统计相同字母连续出现的长度
// 例如字符串 aaaabbbabb 分成 aaaa+bbb+a+bb 四组， 字母a长度列表为[4, 1] 字母b长度列表为[3,2]

// 遍历每个字母对应的长度列表a,把a从大到小排序

// 取出三个特殊子串的方法
// 1. 从最长的特殊子串 a[0] 中取三个长度均为 a[0] - 2 的特殊子串，

// 2. 从最长和次长的特殊子串(a[0], a[1]) 取三个长度一样的特殊字串：
// 	 2.1 如果 a[0] = a[1]，那么可以取出三个长度均为a[0]-1的特殊子串
//	 2.2 如果 a[0] > a[1], 那么可以取出三个长度均为a[1]的特殊子串： 从最长中取两个，从次长中取一个
//	 合并为 min(a[0]-1, a[1])

// 3. 从最长、次长、第三长的特殊子串(a[0],a[1],a[2]) 中各取一个长为 a[2] 的特殊子串。
//	取最大值 max(a[0]-2, min(a[0]-1, a[1]), a[2])

func maximumLength(s string) (ans int) {
	groups := [26][]int{}
	cnt := 0
	for i := range s {
		cnt++
		if i+1 == len(s) || s[i] != s[i+1] {
			groups[s[i]-'a'] = append(groups[s[i]-'a'], cnt) // 连续字符的长度
			cnt = 0
		}
	}

	for _, a := range groups {
		if len(a) == 0 {
			continue
		}
		slices.SortFunc(a, func(a, b int) int { return b - a })
		a = append(a, 0, 0)
		ans = max(ans, a[0]-2, min(a[0]-1, a[1]), a[2])
	}

	if ans == 0 {
		return -1
	}
	return
}
