package main

func backspaceCompare1(s string, t string) bool {
	st1, st2 := []byte{}, []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '#' && len(st1) > 0 {
			st1 = st1[:len(st1)-1]
		} else {
			st1 = append(st1, s[i])
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i] == '#' && len(st2) > 0 {
			st2 = st2[:len(st2)-1]
		} else {
			st2 = append(st2, t[i])
		}
	}
	return string(st1) == string(st2)
}

func backspaceCompare(s string, t string) bool {
	p1, p2 := len(s)-1, len(t)-1
	skipS, skipT := 0, 0
	for p1 >= 0 || p2 >= 0 {
		for p1 >= 0 {
			if s[p1] == '#' {
				skipS++
				p1--
			} else if skipS > 0 {
				skipS--
				p1--
			} else {
				break
			}
		}
		for p2 >= 0 {
			if t[p2] == '#' {
				skipT++
				p2--
			} else if skipT > 0 {
				skipT--
				p2--
			} else {
				break
			}
		}
		if p1 >= 0 && p2 >= 0 {
			if s[p1] != t[p2] {
				return false
			}
		} else if p1 >= 0 || p2 >= 0 {
			return false
		}
		p1--
		p2--
	}
	return true
}
