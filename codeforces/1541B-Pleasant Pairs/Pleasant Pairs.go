package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1541B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var t, n, aj int
	Fscan(in, &t)
	for k := 0; k < t; k++ {
		Fscan(in, &n)
		idx := make([]int, n*2+1)
		ans := 0
		for j := 1; j <= n; j++ {
			Fscan(in, &aj)
			for ai := 1; ai*aj < j*2; ai++ {
				i := idx[ai]
				if i > 0 && ai*aj == i+j {
					ans++
				}
			}
			idx[aj] = j
		}
		Fprintln(out, ans)
	}
}

func main() { CF1541B(os.Stdin, os.Stdout) }
