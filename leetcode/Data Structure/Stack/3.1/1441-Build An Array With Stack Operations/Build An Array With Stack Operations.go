package main

func buildArray1(target []int, n int) (ans []string) {
	cnt := 1
	index := 0
	for cnt <= n {
		ans = append(ans, "Push")
		if cnt == target[index] {
			index++
		} else {
			ans = append(ans, "Pop")
		}
		cnt++
	}
	return
}

func buildArray(target []int, n int) (ans []string) {
	st := []int{}
	for i := 0; i < n; i++ {
		st = append(st, i+1)
		ans = append(ans, "Push")
		if st[len(st)-1] != target[len(st)-1] {
			ans = append(ans, "Pop")
			st = st[:len(st)-1]
		}
		if len(target) == len(st) {
			return
		}
	}
	return
}
