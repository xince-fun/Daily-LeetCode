package main

// (a ^ b) ^ a = a

// (a ^ b) ^ b = b

func decode1(encoded []int, first int) (ans []int) {
	ans = append(ans, first)
	for i, x := range encoded {
		ans = append(ans, x^ans[i])
	}
	return
}

func decode(encoded []int, first int) []int {
	encoded = append([]int{first}, encoded...)
	for i := 1; i < len(encoded); i++ {
		encoded[i] = encoded[i-1] ^ encoded[i]
	}
	return encoded
}
