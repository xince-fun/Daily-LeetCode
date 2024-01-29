package main

func flowerGame(n int, m int) int64 {
	oddCountX := (n + 1) / 2
	evenCountX := n / 2

	oddCountY := (m + 1) / 2
	evenCountY := m / 2

	return int64(oddCountX)*int64(evenCountY) + int64(evenCountX)*int64(oddCountY)
}
