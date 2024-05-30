package main

func removeStars(s string) string {
	st := []rune{}
	for _, c := range s {
		st = append(st, c)
		if c == '*' {
			st = st[:len(st)-1]
			st = st[:len(st)-1]
		}
	}
	return string(st)
}
