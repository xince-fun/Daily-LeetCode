package main

import (
	"fmt"
)

func removeAlmostEqualCharacters(word string) int {
	operations := 0
	// 遍历字符串，直到倒数第二个字符
	for i := 0; i < len(word)-1; i++ {
		// 当前字符与下一个字符近似相等
		if isAlmostEqual(word[i], word[i+1]) {
			operations++
			// 更改当前字符后，跳过下一个字符
			i++
		}
	}
	return operations
}

// 判断字符是否近似相等（相同或者在字母表中相邻）
func isAlmostEqual(a, b byte) bool {
	return a == b || abs(int(a)-int(b)) == 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	word := "abddez"
	fmt.Println(removeAlmostEqualCharacters(word))
}
