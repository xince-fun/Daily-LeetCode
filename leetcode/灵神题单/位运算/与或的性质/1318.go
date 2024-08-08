package main

func minFlips(a int, b int, c int) (ans int) {
	for i := 0; i < 31; i++ {
		bit_a, bit_b, bit_c := (a>>i)&1, (b>>i)&1, (c>>i)&1
		if bit_c == 1 {
			if bit_a+bit_b == 0 {
				ans++
			}
		} else {
			ans += bit_a + bit_b
		}
	}
	return
}
