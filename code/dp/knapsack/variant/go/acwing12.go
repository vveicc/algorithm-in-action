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

	var n, m int
	Fscan(in, &n, &m)
	items := make([][2]int, n)
	for i := range items {
		Fscan(in, &items[i][0], &items[i][1])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// 倒序转移最大总价值
	for i := n - 1; i >= 0; i-- {
		v, w := items[i][0], items[i][1]
		for j := 1; j <= m; j++ {
			if dp[i][j] = dp[i+1][j]; j >= v {
				dp[i][j] = max(dp[i][j], dp[i+1][j-v]+w)
			}
		}
	}

	// 正序输出字典序最小的方案
	left := m
	for i := 0; i < n && left > 0; i++ {
		v, w := items[i][0], items[i][1]
		if left >= v && dp[i][left] == dp[i+1][left-v]+w {
			Fprint(out, i+1, " ")
			left -= v
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
