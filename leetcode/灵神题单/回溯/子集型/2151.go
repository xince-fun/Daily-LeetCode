package main

/*
有 n 个人，每个人可能是好人，也可能是坏人，
因此这 n 个人的好坏有 2^n 种情况。

二进制枚举好人（一共 2^n 种情况），
然后根据好人的陈述来判断合法性：

因为好人只说真话，那么：
  - 好人对好人的陈述只能是 1 或 2，即不能是 0
  - 好人对坏人的陈述只能是 0 或 2，即不能是 1
*/
func maximumGood(statements [][]int) (ans int) {
next:
	for i := 1; i < 1<<len(statements); i++ {
		cnt := 0 // i 中好人个数
		for j, row := range statements {
			if i>>j&1 == 1 { // 枚举 i 中的好人 j
				for k, st := range row { // 枚举 j 的所有陈述 st
					if st < 2 && st != i>>k&1 { // 该陈述与实际情况矛盾
						continue next
					}
				}
				cnt++
			}
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return
}
