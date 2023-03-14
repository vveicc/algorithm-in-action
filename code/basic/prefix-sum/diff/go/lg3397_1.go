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

	var n, m, x1, y1, x2, y2 int
	Fscan(in, &n, &m)
	diff := make([][]int, n+1)
	for i := range diff {
		diff[i] = make([]int, n+2)
	}
	for ; m > 0; m-- {
		Fscan(in, &x1, &y1, &x2, &y2)
		for i := x1; i <= x2; i++ {
			diff[i][y1]++
			diff[i][y2+1]--
		}
	}

	for i := 1; i <= n; i++ {
		x := 0
		for j := 1; j <= n; j++ {
			x += diff[i][j]
			Fprint(out, x, " ")
		}
		Fprintln(out)
	}
}
