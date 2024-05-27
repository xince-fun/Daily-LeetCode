package main

func shiftingLetters(s string, shifts [][]int) string {
	mn := 50002
	mx := 0

	for _, t := range shifts {
		mn = min(mn, t[0])
		mx = max(mx, t[1])
	}

	d := make([]int, mx-mn+2)
	for _, t := range shifts {
		if t[2] == 0 {
			d[t[0]-mn]--
			d[t[1]-mn+1]++
		} else if t[2] == 1 {
			d[t[0]-mn]++
			d[t[1]-mn+1]--
		}
	}
	sum := 0
	ans := []byte{}
	i := 0
	for i < mn {
		ans = append(ans, s[i])
		i++
	}
	for _, v := range d[:len(d)-1] {
		sum += v
		char := (int(s[i])+sum%26+26-'a')%26 + 'a'
		i++
		ans = append(ans, byte(char))
	}
	for i < len(s) {
		ans = append(ans, s[i])
		i++
	}
	return string(ans)
}

func shiftingLetters1(s string, shifts [][]int) string {
	diff := make([]int, len(s)+1)
	for _, p := range shifts {
		x := p[2]*2 - 1 // 0 和 1 变成 -1 和 1
		diff[p[0]] += x
		diff[p[1]+1] -= x
	}
	t, shift := []byte(s), 0
	for i, c := range t {
		shift = (shift+diff[i])%26 + 26
		t[i] = (c-'a'+byte(shift))%26 + 'a'
	}
	return string(t)
}
