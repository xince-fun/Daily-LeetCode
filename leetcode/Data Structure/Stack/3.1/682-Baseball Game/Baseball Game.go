package main

import "strconv"

func calPoints(operations []string) (ans int) {
	st := []int{}
	for _, op := range operations {
		switch op {
		case "C":
			st = st[:len(st)-1]
		case "D":
			new := st[len(st)-1] * 2
			st = append(st, new)
		case "+":
			new := st[len(st)-1] + st[len(st)-2]
			st = append(st, new)
		default:
			num, _ := strconv.Atoi(op)
			st = append(st, num)
		}
	}
	for _, x := range st {
		ans += x
	}
	return
}
