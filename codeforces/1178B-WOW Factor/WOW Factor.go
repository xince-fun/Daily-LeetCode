package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 计算每个o的左边有多少个v，右边有多少个v

func CF1178B1(_r io.Reader, out io.Writer) {
	var str string
	Fscan(bufio.NewReader(_r), &str)
	var pre, suf, ans int
	n := len(str)
	for i := 1; i < n-1; i++ {
		if str[i] == 'v' && str[i+1] == 'v' {
			suf++
		}
	}
	for i := 1; i < n-2; i++ {
		if str[i] == 'o' {
			ans += pre * suf
		} else {
			if str[i-1] == 'v' {
				pre++
			}
			if str[i+1] == 'v' {
				suf--
			}
		}
	}
	Print(ans)
}

// 状态机dp

// f_{i,0}表示考虑前i个字母，能得到的w子序列的个数
// f_{i,1}表示考虑前i个字母，能得到的wo子序列的个数
// f_{i,2}表示考虑前i个字母，能得到的wow子序列的个数

// 分类讨论
// 如果 s_{i} = 'o',那么 f_{i,1} = f_{i-1,1} + f_{i-1, 0} 表示分别不选或选 s_i
// 如果 s_{i} = 'v' 且 s_{i-1} = 'v' ,那么 f_{i,0} = f_{i-1,0} + 1, f_{i,2} = f_{i-1,2} + f_{i-1,1}

func CF1178B(_r io.Reader, out io.Writer) {
	var str string
	Fscan(bufio.NewReader(_r), &str)
	var f0, f1, f2 int
	for i := 1; i < len(str); i++ {
		if str[i] == 'o' {
			f1 += f0
		} else if str[i-1] == 'v' {
			f2 += f1
			f0++
		}
	}
	Print(f2)
}

func main() { CF1178B(os.Stdin, os.Stdout) }
