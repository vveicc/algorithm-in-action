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

	var n int
	Fscan(in, &n)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	b := make([]int64, n)
	for i := range b {
		Fscan(in, &b[i])
		b[i]--
	}
	vis := make([]bool, n)
	for i := n - 1; i >= 0; i-- {
		k := b[i]
		b[i], vis[k] = 0, true
		for x := range a {
			for y := range a {
				a[x][y] = min(a[x][y], a[x][k]+a[k][y])
				if vis[x] && vis[y] {
					b[i] += int64(a[x][y])
				}
			}
		}
	}
	for _, x := range b {
		Fprintf(out, "%d ", x)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
