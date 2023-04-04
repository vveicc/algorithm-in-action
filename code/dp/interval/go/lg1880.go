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
	a := make([]int, n)
	s := make([]int, n<<1+1)
	for i := range a {
		Fscan(in, &a[i])
		s[i+1] = s[i] + a[i]
	}
	for i, x := range a {
		i += n
		s[i+1] = s[i] + x
	}
	mn := make([][]int, n<<1)
	mx := make([][]int, n<<1)
	for i := n<<1 - 1; i >= 0; i-- {
		mn[i] = make([]int, n<<1)
		mx[i] = make([]int, n<<1)
		for j := i + 1; j < i+n && j < n<<1; j++ {
			mn[i][j] = 1e6
			for k := i; k < j; k++ {
				mn[i][j] = min(mn[i][j], mn[i][k]+mn[k+1][j])
				mx[i][j] = max(mx[i][j], mx[i][k]+mx[k+1][j])
			}
			// 合并成一堆
			score := s[j+1] - s[i]
			mn[i][j] += score
			mx[i][j] += score
		}
	}
	x := mn[0][n-1]
	y := mx[0][n-1]
	for i := 1; i < n; i++ {
		x = min(x, mn[i][i+n-1])
		y = max(y, mx[i][i+n-1])
	}
	Fprintln(out, x)
	Fprintln(out, y)
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
