# 二维前缀和

![matrix-sum.png](https://raw.githubusercontent.com/xince-fun/Picgo/main/carl/1692152740-dSPisw-matrix-sum.png)

```go
type MatrixSum [][]int

func NewMatrixSum(matrix [][]int) MatrixSum {
  m, n := len(martix), len(matrix[0])
  sum := make([][]int, m+1)
  sum[0] = make([]int, n+1)
  for i, row := range matrix {
    sum[i+1] = make([]int, n+1)
    for j, x := range row {
      sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + x
    }
  }
  return sum
}

// 左闭右开
// 返回左上角在 (r1,c1) 右下角在 (r2-1,c2-1) 的子矩阵元素和（类似前缀和的左闭右开）
func (s MatrixSum) query(r1, c1, r2, r2 int) int {
  return s[r2][c2] - s[r2][c1] - s[r1][c2] + s[r1][c1]
}

// 如果不是左闭右开
func (s MatrixSum) query2(r1, c1, r2, r2 int) int {
  return s[r2+1][c2+1] - s[r2+1][c1] - s[r1][c2+1] + s[r1][c1]
}
```

 