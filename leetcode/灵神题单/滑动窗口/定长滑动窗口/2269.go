package main

import "strconv"

// 4124778

// 4
// 40+1=41
// 410+2=412
// 4120+4=4124
// 4124 => 124

func divisorSubstrings1(num int, k int) (ans int) {
	str := strconv.Itoa(num)
	sum := 0
	for i, in := range str {
		sum = sum*10 + int(in-'0')
		if i < k-1 {
			continue
		}
		if num/sum == 0 {
			ans++
		}
		out := str[i-k+1]
		sum -= int(out-'0') * pow(10, k)
	}
	return ans
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x
		}
		x = x * x
	}
	return res
}

func divisorSubstrings2(num int, k int) (ans int) {
	str := strconv.Itoa(num)
	for i := k; i <= len(str); i++ {
		v, _ := strconv.Atoi(str[i-k : i])
		if v > 0 && num%v == 0 {
			ans++
		}
	}
	return ans
}

func divisorSubstrings(num int, k int) (ans int) {
	str := strconv.Itoa(num)
	for i := range str {
		if i < k-1 {
			continue
		}
		v, _ := strconv.Atoi(str[i-k+1 : i+1])
		if v > 0 && num%v == 0 {
			ans++
		}
	}
	return ans
}
