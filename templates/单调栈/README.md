# 单调栈

```go

func () {
    st := []int{}
    for i := n - 1; i >= 0; i-- {
        x := nums[i]
        // 从左向右看是递减 从右往左看是递增
        for len(st) != 0 && x <= nums[len(st)-1] {
            st = st[:len(st)-1]
        }
    }
}

```