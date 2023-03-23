package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]string, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		Fscan(in, &s)

		fi2 := func(t string, i int) int { return int(t[i]&15)*10 + int(t[i+1]&15) }
		fi3 := func(t string, i int) int { return int(t[i]&15)*100 + int(t[i+1]&15)*10 + int(t[i+2]&15) }
		fs2 := func(i int) int { return fi2(s, i) }
		fs3 := func(i int) int { return fi3(s, i) }

		// 预处理a中所有长度为2和3的子串的位置
		f2 := make([][3]int, 100)
		f3 := make([][3]int, 1000)
		for i, t := range a {
			if m > 1 {
				f2[fi2(t, m-2)] = [3]int{m - 1, m, i + 1}
			}
			for j := m - 3; j >= 0; j-- {
				f2[fi2(t, j)] = [3]int{j + 1, j + 2, i + 1}
				f3[fi3(t, j)] = [3]int{j + 1, j + 3, i + 1}
			}
		}

		// f[i]表示s[i:]是否可以划分
		f := make([]bool, m+1)
		f[m] = true
		if m > 1 {
			f[m-2] = f2[fs2(m-2)][0] != 0
		}
		for i := m - 3; i >= 0; i-- {
			f[i] = (f[i+2] && f2[fs2(i)][0] != 0) || (f[i+3] && f3[fs3(i)][0] != 0)
		}

		if !f[0] {
			Fprintln(out, -1)
		} else {
			var arr [][3]int
			for i := 0; i != m; {
				if x := fs2(i); f2[x][0] != 0 && f[i+2] {
					arr = append(arr, f2[x])
					i += 2
				} else {
					arr = append(arr, f3[fs3(i)])
					i += 3
				}
			}
			Fprintln(out, len(arr))
			for _, i := range arr {
				Fprintln(out, i[0], i[1], i[2])
			}
		}
	}
}
