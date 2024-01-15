package main

func addMinimum(word string) int {
	// 统计所有 a 出现的次数
	t := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] >= word[i] {
			t++
		}
	}
	return t*3 - len(word)
}

func addMinimum1(word string) int {
	ans := int(word[0]) - int(word[len(word)-1]) + 2
	for i := 1; i < len(word); i++ {
		ans += (int(word[i]) - int(word[i-1]) + 2) % 3
	}
	return ans
}

func addMinimum2(word string) (ans int) {
	s := "abc"
	n := len(word)
	for i, j := 0, 0; j < n; i = (i + 1) % 3 {
		if word[j] != s[i] {
			ans++
		} else {
			j++
		}
	}
	if word[n-1] == 'b' {
		ans += 1
	}
	if word[n-1] == 'a' {
		ans += 2
	}
	return
}
