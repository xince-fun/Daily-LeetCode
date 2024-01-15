package main

import (
	"fmt"
)

func repeatLimitedString(s string, repeatLimit int) string {
	ans := []byte{}
	arr := [26]int{}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	for i, j := 25, 24; i >= 0; i-- {
		j = min(j, i-1)
		for {
			for k := min(arr[i], repeatLimit); k > 0; k-- {
				ans = append(ans, byte(i)+'a')
				arr[i]--
			}
			if arr[i] == 0 {
				break
			}
			for j >= 0 && arr[j] == 0 {
				j--
			}
			if j < 0 {
				break
			}
			ans = append(ans, byte(j)+'a')
			arr[j]--
		}
	}
	return string(ans)
}

func main() {
	s := "cczazcc"
	repeatLimit := 3
	fmt.Println(repeatLimitedString(s, repeatLimit))
}
