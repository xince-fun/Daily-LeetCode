# 子集型回溯

// 子集型回溯
// 每个元素都可以选/不选

## 模板1


> // 回溯三问
> // 1. 当前操作？枚举第i个数选/不选
> // 2. 子问题？ 从下标>=i的数字中构造子集
> // 3. 下一个子问题？ 从下标>=i+1的数字中构造子集
> // dfs(i) -> dfs(i+1)

```go
func subsets1(nums []int) (ans [][]int) {
	n := len(nums)

	path := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(i + 1) // 不选
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(0)
	return
}
```

## 模板2

> // 回溯三问
> // 1. 当前操作？构造一个下标 j >= i的数字，加入path
> // 2. 子问题？ 从下标 >= i的数字中构造子集
> // 3. 下一个子问题？ 从下标>=j+1的数字中构造子集
> // dfs(i) -> dfs(i+1)
> // dfs(i) -> dfs(i+2)

```go
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	path := []int{}

	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, append([]int(nil), path...))
		if i == n {
			return
		}
		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
```