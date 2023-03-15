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
	arr := make([][2]int, n)
	for i := range arr {
		Fscan(in, &arr[i][0], &arr[i][1])
	}

	dp := make([]int, m+1)
	for _, item := range arr { // 枚举物品
		v, w := item[0], item[1]
		for j := m; j >= v; j-- { // 枚举背包容量
			dp[j] = max(dp[j], dp[j-v]+w)
		}
	}
	Fprintln(out, dp[m])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
