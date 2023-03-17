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
	var chairs, peoples []int
	for i := 0; i < n; i++ {
		if Fscan(in, &x); x == 0 {
			chairs = append(chairs, i)
		} else {
			peoples = append(peoples, i)
		}
	}

	n = len(peoples)
	dp := make([]int, n+1)
	for j := range dp {
		dp[j] = 1e8
	}
	dp[0] = 0
	for i, chair := range chairs {
		for j := min(n-1, i); j >= 0; j-- {
			dp[j+1] = min(dp[j+1], dp[j]+abs(chair-peoples[j]))
		}
	}
	Fprintln(out, dp[n])
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
