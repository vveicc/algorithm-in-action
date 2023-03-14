package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var s []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &s)
		// 排序
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		var a [2][]byte
		for _, b := range s {
			a[b&1] = append(a[b&1], b) // 按序奇偶分组
		}
		x, y := a[0], a[1]
		if m, n := len(x), len(y); m == 0 {
			Fprintf(out, "%s\n", y)
		} else if n == 0 {
			Fprintf(out, "%s\n", x)
		} else if abs(int(x[m-1])-int(y[0])) != 1 {
			Fprintf(out, "%s%s\n", x, y) // x+y
		} else if abs(int(y[n-1])-int(x[0])) != 1 {
			Fprintf(out, "%s%s\n", y, x) // y+x
		} else {
			// 当且仅当字符串s仅由2个或3个相邻字母组成时，无法完成构造
			Fprintln(out, "No answer")
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
