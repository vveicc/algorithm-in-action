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
	diff := make([][]int, n+2)
	for i := range diff {
		diff[i] = make([]int, n+2)
	}
	for ; m > 0; m-- {
		Fscan(in, &x1, &y1, &x2, &y2)
		diff[x1][y1]++
		diff[x1][y2+1]--
		diff[x2+1][y1]--
		diff[x2+1][y2+1]++
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			diff[i][j] += diff[i-1][j] + diff[i][j-1] - diff[i-1][j-1]
			Fprint(out, diff[i][j], " ")
		}
		Fprintln(out)
	}
}
