package main

func calculate(s string) int {

	var cal func(*[]byte) int

	cal = func(ss *[]byte) int {
		st := []int{}
		num := 0
		sign := byte('+')
		for len(*ss) > 0 {
			c := (*ss)[0]
			(*ss) = (*ss)[1:] // pop left
			if c >= '0' && c <= '9' {
				num = num*10 + int(c-'0')
			}
			if c == '(' {
				num = cal(ss)
			}

			if (!(c >= '0' && c <= '9') && c != ' ') || len(*ss) == 0 {
				switch sign {
				case '+':
					st = append(st, num)
				case '-':
					st = append(st, -num)
				case '*':
					st[len(st)-1] *= num
				case '/':
					st[len(st)-1] /= num
				}
				num = 0
				sign = c
			}

			if c == ')' {
				break
			}
		}

		sum := 0
		for _, x := range st {
			sum += x
		}
		return sum
	}
	ss := make([]byte, len(s))
	copy(ss, s)
	return cal(&ss)
}
