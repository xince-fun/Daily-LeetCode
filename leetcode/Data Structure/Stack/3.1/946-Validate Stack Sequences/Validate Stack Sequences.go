package main

func validateStackSequences1(pushed []int, popped []int) bool {
	p := 0
	st := []int{}
	for i := 0; i < len(pushed); i++ {
		st = append(st, pushed[i])
		for len(st) > 0 && st[len(st)-1] == popped[p] {
			st = st[:len(st)-1]
			p++
		}
	}
	return len(st) == 0
}

func validateStackSequences(pushed []int, popped []int) bool {
	st, p := []int{}, 0
	for _, val := range pushed {
		st = append(st, val)
		for len(st) > 0 && st[len(st)-1] == popped[p] {
			p++
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}
