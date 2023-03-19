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

	var n, m, c int
	Fscan(in, &n, &m)
	dp := make([]int, m+1)
	for ; n > 0; n-- { // 枚举分组
		Fscan(in, &c)
		lower := 100
		items := make([][2]int, c)
		for i := range items {
			Fscan(in, &items[i][0], &items[i][1])
			lower = min(lower, items[i][0])
		}
		for i := m; i >= lower; i-- { // 枚举背包容量
			for _, item := range items { // 枚举分组内的每一件物品
				if v, w := item[0], item[1]; i >= v {
					dp[i] = max(dp[i], dp[i-v]+w)
				}
			}
		}
	}
	Fprintln(out, dp[m])
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
