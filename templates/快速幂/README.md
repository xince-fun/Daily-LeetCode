# 快速幂

```go
const mod int = 1e9 + 7

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

```go
func myPow(x float64, n int) float64 {
	var res float64 = 1
    if n < 0 {
        x, n = 1 / x, -n
    }
    for ; n > 0; n /= 2 {
        if n&1 > 0 {
            res = res * x
        }
        x = x * x
    }
    return res
}
```