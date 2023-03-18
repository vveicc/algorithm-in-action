package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	Fscan(in, &n, &m)
	items := make([][2]int, n)
	for i := range items {
		Fscan(in, &items[i][0], &items[i][1])
	}

	dp := make([]int, m+1)
	for _, item := range items { // 枚举物品
		v, w := item[0], item[1]
		for j := v; j <= m; j++ { // 枚举背包容量
			dp[j] = max(dp[j], dp[j-v]+w)
		}
	}
	Println(dp[m])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
