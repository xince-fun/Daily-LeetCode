package main

import "fmt"

func splitMessage(message string, limit int) []string {
	for i, cap, tailLen := 1, 0, 0; ; i++ {
		if i < 10 {
			tailLen = 5 // 结尾的长度
		} else if i < 100 {
			if i == 10 {
				cap -= 9
			} // 前面的结尾的长度都 +1，那么容量就要减小
			tailLen = 7
		} else if i < 1000 {
			if i == 100 {
				cap -= 99
			}
			tailLen = 9
		} else {
			if i == 1000 {
				cap -= 999
			}
			tailLen = 11
		}
		if tailLen >= limit {
			return nil
		} // cap 无法增大，寄
		cap += limit - tailLen
		if cap < len(message) {
			continue
		} // 容量没有达到，继续枚举

		ans := make([]string, i)
		for j := range ans {
			tail := fmt.Sprintf("<%d/%d>", j+1, i)
			if j == i-1 {
				ans[j] = message + tail
			} else {
				m := limit - len(tail)
				ans[j] = message[:m] + tail
				message = message[m:]
			}
		}
		return ans
	}

}
