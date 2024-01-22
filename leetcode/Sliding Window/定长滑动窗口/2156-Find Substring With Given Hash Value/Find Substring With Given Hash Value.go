package main

import (
	"fmt"
)

func subStrHash(s string, power int, modulo int, k int, hashValue int) (ans string) {
	hash, p := 0, 1
	i, n := len(s)-1, len(s)
	for ; i >= n-k; i-- {
		hash = (hash*power + int(s[i]&31)) % modulo // 用秦九韶算法计算 s[n-k:] 的哈希值
		p = p * power % modulo
	}
	if hash == hashValue {
		ans = s[n-k:]
	}
	for ; i >= 0; i-- {
		hash = (hash*power + int(s[i]&31) + modulo - p*int(s[i+k]&31)%modulo) % modulo
		if hash == hashValue {
			ans = s[i : i+k]
		}
	}
	return
}

func main() {
	s := "xqgfatvtlwnnkxipmipcpqwbxihxblaplpfckvxtihonijhtezdnkjmmk"
	power := 22
	module := 51
	k := 41
	hashValue := 9
	fmt.Println(subStrHash(s, power, module, k, hashValue))
}
