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

	var N, C, M int
	Fscan(in, &N, &C, &M)
	items := make([][3]int, N)
	for i := range items {
		Fscan(in, &items[i][0], &items[i][1], &items[i][2])
	}

	dp := make([][]int, C+1)
	for i := range dp {
		dp[i] = make([]int, M+1)
	}
	for _, item := range items { // 枚举物品
		v, m, w := item[0], item[1], item[2]
		for i := C; i >= v; i-- { // 枚举第一维度费用
			for j := M; j >= m; j-- { // 枚举第二维度费用
				dp[i][j] = max(dp[i][j], dp[i-v][j-m]+w)
			}
		}
	}
	Fprintln(out, dp[C][M])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
