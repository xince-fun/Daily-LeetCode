package main

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxLenth := 0
	maxArea := -1

	for _, d := range dimensions {
		x := d[0]*d[0] + d[1]*d[1]
		if x > maxLenth {
			maxLenth = x
			maxArea = d[0] * d[1]
		} else if x == maxLenth && d[0]*d[1] > maxArea {
			maxArea = d[0] * d[1]
		}
	}
	return maxArea
}
