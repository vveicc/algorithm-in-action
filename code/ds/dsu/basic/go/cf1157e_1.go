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

	var n, x int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		b[x]++
	}

	// 并查集
	fa := make([]int, n+1)
	for i, next := n, 0; i >= 0; i-- {
		if b[i] != 0 {
			next = i
		}
		fa[i] = next
	}

	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x]) // 路径压缩
		}
		return fa[x]
	}
	union := func(f, t int) {
		if f, t = find(f), find(t); f != t {
			fa[f] = t
		}
	}

	for _, x = range a {
		y := find(n - x)
		if b[y]--; b[y] == 0 {
			union(y, y+1) // 指向下一个未使用的数字
		}
		Fprint(out, (x+y)%n, " ")
	}
}
