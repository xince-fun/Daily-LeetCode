package main

import "strconv"

func divisorSubstrings1(num int, k int) (ans int) {
	str := strconv.Itoa(num)
	left := 0
	for right := range str {
		for right-left+1 > k {
			left++
		}
		if right-left+1 == k {
			s := str[left : right+1]
			n, _ := strconv.Atoi(s)
			if n != 0 && num%n == 0 {
				ans++
			}
		}
	}
	return
}

func divisorSubstrings(num int, k int) (ans int) {
	str := strconv.Itoa(num)
	for i := k; i <= len(str); i++ {
		v, _ := strconv.Atoi(str[i-k : i])
		if v > 0 && num%v == 0 {
			ans++
		}
	}
	return
}
