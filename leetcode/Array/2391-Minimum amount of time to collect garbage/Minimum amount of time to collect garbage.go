package main

// func garbageCollection(garbage []string, travel []int) int {
// 	ans := 0
// 	// 提前记录每个房子位置的垃圾 0为M，1为P, 2为G
// 	// 提前记录每个车的终止位置
// 	// stop0, stop1, stop2 := 0, 0, 0
// 	// g := make([][3]int, len(garbage))
// 	// for i, gab := range garbage {
// 	// 	g0, g1, g2 := 0, 0, 0
// 	// 	for _, ga := range gab {
// 	// 		switch ga {
// 	// 		case 'M':
// 	// 			g0++
// 	// 		case 'P':
// 	// 			g1++
// 	// 		case 'G':
// 	// 			g2++
// 	// 		}
// 	// 	}
// 	// 	if g0 > 0 {
// 	// 		ans += g0
// 	// 		// ans += travel[i]
// 	// 		// if i > 0 {
// 	// 		// 	ans += travel[i-1]
// 	// 		// }
// 	// 		// stop0 = i
// 	// 		// g[i][0] += g0
// 	// 	}
// 	// 	if g1 > 0 {
// 	// 		ans += g1
// 	// 		// if i > 0 {
// 	// 		// 	ans += travel[i-1]
// 	// 		// }
// 	// 		// stop1 = i
// 	// 		// g[i][1] += g1
// 	// 	}
// 	// 	if g2 > 0 {
// 	// 		// stop2 = i
// 	// 		ans += g2
// 	// 		// if i > 0 {
// 	// 		// 	ans += travel[i-1]
// 	// 		// }
// 	// 		// g[i][2] += g2
// 	// 	}
// 	// }
// 	stop0, stop1, stop2 := false, false, false
// 	for i := len(garbage) - 1; i >= 0; i-- {
// 		gab := garbage[i]
// 		// 如果要到这里
// 		ans += len(gab)
// 		if !stop0 || !stop1 || !stop2 {
// 			for _, g := range gab {
// 				if !stop0 && g == 'M' {
// 					stop0 = true
// 					for _, t := range travel[:i] {
// 						ans += t
// 					}
// 				}
// 				if !stop1 && g == 'P' {
// 					stop1 = true
// 					for _, t := range travel[:i] {
// 						ans += t
// 					}
// 				}
// 				if !stop2 && g == 'G' {
// 					stop2 = true
// 					for _, t := range travel[:i] {
// 						ans += t
// 					}
// 				}
// 			}
// 		}

// 	}
// 	return ans
// }

func garbageCollection(garbage []string, travel []int) int {
	sum := len(garbage[0])
	metM, metP, metG := 0, 0, 0
	for i := len(garbage) - 1; i > 0; i-- {
		for _, g := range garbage[i] {
			if g == 'M' {
				metM = 1
			}
			if g == 'P' {
				metP = 1
			}
			if g == 'G' {
				metG = 1
			}
		}
		sum += len(garbage[i]) + (metM+metP+metG)*travel[i-1]
	}
	return sum
}
