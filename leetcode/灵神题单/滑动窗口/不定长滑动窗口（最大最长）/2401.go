package main

// 3=> 0011
// 8=> 1000

/*
子数组中任意两个数按位与均为 0，意味着任意两个数对应的集合的交集为空（见 从集合论到位运算，常见位运算技巧分类总结）。

这意味着子数组中的从低到高的第 i 个比特位上，至多有一个比特 1，其余均为比特 0。例如子数组不可能有两个奇数（最低位为 1)，否则这两个数按位与是大于 0 的。

根据鸽巢原理（抽屉原理），在本题数据范围下，优雅子数组的长度不会超过 30。例如子数组为 [2^0,2^1,2^2,⋯,2^29]，我们无法再加入一个数 x，使 x 与子数组中的任何一个数按位与均为 0。

既然长度不会超过 30，直接暴力枚举子数组的右端点 i 即可。

代码实现时，可以把在子数组中的元素按位或起来（求并集），这样可以 O(1) 判断当前元素是否与前面的元素按位与的结果为 0（交集为空）。
*/

// 最长优雅子数组的长度不超过30
func longestNiceSubarray1(nums []int) (ans int) {
	for i, or := range nums {
		j := i - 1
		for ; j >= 0 && or&nums[j] == 0; j-- {
			or |= nums[j]
		}
		ans = max(ans, i-j)
	}
	return
}

func longestNiceSubarray(nums []int) (ans int) {
	left, or := 0, 0
	for right, x := range nums {
		for or&x > 0 { // 有交集
			or ^= nums[left] // 从 or中去掉
			left++
		}
		or |= x
		ans = max(ans, right-left+1)
	}
	return
}
