package main

// func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
// 	if a == e || b == f {
// 		// 检查是否有棋子在车和黑皇后之间
// 		if !((a == c && ((b < f && f < d) || (d < f && f < b))) ||
// 			(b == d && ((a < e && e < c) || (c < e && e < a)))) {
// 			return 1 // 没有棋子阻挡，车可以直接捕获黑皇后
// 		}
// 	}

// 	// 如果象在与黑皇后同一对角线上
// 	if abs(c-e) == abs(d-f) {
// 		// 检查是否有棋子在象和黑皇后之间
// 		blocked := false
// 		delta := abs(c - e)
// 		for i := 1; i < delta; i++ {
// 			if (a == c+i && b == d+i) || (a == c+i && b == d-i) ||
// 				(a == c-i && b == d+i) || (a == c-i && b == d-i) {
// 				blocked = true
// 				break
// 			}
// 		}
// 		if !blocked {
// 			return 1 // 没有棋子阻挡，象可以直接捕获黑皇后
// 		}
// 	}

// 	// 如果车或者象不能一步捕获黑皇后，那么至少需要两步
// 	// 第一步是移动车或者象到能够下一步捕获黑皇后的位置
// 	// 第二步是捕获黑皇后
// 	return 2
// }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
白车可以直接攻击到黑后 -> 1
百象可以直接攻击到黑后 -> 1

一定是 2

*/

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	notInside := func(left, mid, right int) bool {
		return !(min(left, right) < mid && mid < max(left, right))
	}

	if a == e && (c != e || notInside(b, d, f)) ||
		b == f && (d != f || notInside(a, c, e)) ||
		c+d == e+f && (a+b != e+f || notInside(c, a, e)) ||
		c-d == e-f && (a-b != e-f || notInside(c, a, e)) {
		return 1
	}
	return 2
}
