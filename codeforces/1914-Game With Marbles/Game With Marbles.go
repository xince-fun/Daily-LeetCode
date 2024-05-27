package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// S = sum a_i - sum b_i
// if Bob made all steps: then S = 0 - sum (b_i - 1)
// each Alice step: S += (a_i - 1) + (b_i - 1) i. e. the bigger (a_i + b_i) the better

func CF1904E2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var t int
	Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		Fscan(in, &n)
		a, b := make([]int, n), make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for i := range b {
			Fscan(in, &b[i])
		}
		type pair struct{ x, y int }
		pairs := make([]pair, n)
		for i, x := range a {
			pairs[i] = pair{x, b[i]}
		}
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].x+pairs[i].y > pairs[j].x+pairs[j].y
		})
		diff := 0
		for i, p := range pairs {
			if i%2 == 0 {
				diff += p.x - 1
			} else {
				diff -= p.y - 1
			}
		}
		Fprintln(out, diff)
	}
}

func main() { CF1904E2(os.Stdin, os.Stdout) }
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
