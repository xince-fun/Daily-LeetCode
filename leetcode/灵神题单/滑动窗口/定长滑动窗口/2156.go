package main

func subStrHash(s string, power, mod, k, hashValue int) (ans string) {
	n := len(s)
	// 用秦九韶算法计算 s[n-k:] 的哈希值，同时计算 pk=power^k
	hash, pk := 0, 1
	for i := n - 1; i >= n-k; i-- {
		hash = (hash*power + int(s[i]&31)) % mod
		pk = pk * power % mod
	}
	if hash == hashValue {
		ans = s[n-k:]
	}
	// 向左滑窗
	for i := n - k - 1; i >= 0; i-- {
		// 计算新的哈希值，注意 +mod 防止计算出负数
		hash = (hash*power + int(s[i]&31) - pk*int(s[i+k]&31)%mod + mod) % mod
		if hash == hashValue {
			ans = s[i : i+k]
		}
	}
	return
}
