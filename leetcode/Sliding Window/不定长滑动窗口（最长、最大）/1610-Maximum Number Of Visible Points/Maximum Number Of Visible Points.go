package main

import (
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	eps := 1e-9
	x, y := location[0], location[1]
	cnt := 0
	arr := []float64{}
	t := float64(angle) * math.Pi / 180.0
	for _, p := range points {
		a, b := p[0], p[1]
		if a == x && b == y {
			cnt++
			continue
		}
		arr = append(arr, math.Atan2(float64(b)-float64(y), float64(a)-float64(x))+math.Pi)
	}
	sort.Float64s(arr)
	n := len(arr)
	maxN := 0
	for i := 0; i < n; i++ {
		arr = append(arr, arr[i]+2*math.Pi)
	}
	for i, j := 0, 0; j < 2*n; j++ {
		for i < j && arr[j]-arr[i] > t+eps {
			i++
		}
		maxN = max(maxN, j-i+1)
	}
	return cnt + maxN
}
