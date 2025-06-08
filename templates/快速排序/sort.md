## 快速排序

```go
func sort(nums []int) {
	n := len(nums)
	quickSort(nums, 0, n-1)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}

```

```txt
时间复杂度 O(nlogn)
```

## 快速选择

```go 
nums := []int{5, 4, 2, 6, 1}
k : 3
```

返回第K大的数，就是返回第 n - K 小的数

```go
func topK(nums []int, k int) int {
    n := len(nums)
    return quickSelect(nums, 0, n - 1, n - k)
}

func quickSelect(nums []int, l, r, k int) int {
    if l >= r {
        return nums[k]
    }

    partition := nums[l]
    i := l - 1
    j := r + 1
    for i < j {
        for i++; nums[i] < partition; i++ {}
        for j--; nums[j] > partition; j-- {}
        if i < j {
            nums[i], nums[j] = nums[j], nums[i]
        }
    }

    if k <= j {
        return quickSelect(nums, l, j, k)
    } else {
        return quickSelect(nums, j + 1, r, k)
    }
}

```