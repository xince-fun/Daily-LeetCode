package main

// 从答案的角度来看

// 1. 当前操作? 构建一个下标 j >= i的数字
// 2. 子问题？ 从下标 >= i 中构造子集
// 3. 下一个子问题？ 从下表 >= j+1的数字中构造

func isAdditiveNumber(num string) bool {
	n := len(num)

	var dfs func(i int)
	dfs = func(i int) {

	}
}
