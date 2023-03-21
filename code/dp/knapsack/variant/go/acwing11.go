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

	const mod = 1e9 + 7

	var n, m int
	Fscan(in, &n, &m)
	items := make([][2]int, n)
	for i := range items {
		Fscan(in, &items[i][0], &items[i][1])
	}

	dp := make([]int, m+1)
	// cnt[i]表示容量为i的背包装入总价值为dp[i]的物品的方案数
	cnt := make([]int, m+1)
	for i := range cnt {
		cnt[i] = 1 // 初始化为1，表示什么都不装的方案
	}
	for _, item := range items {
		v, w := item[0], item[1]
		for i := m; i >= v; i-- {
			if x := dp[i-v] + w; x > dp[i] {
				dp[i], cnt[i] = x, cnt[i-v]
			} else if x == dp[i] {
				cnt[i] = (cnt[i] + cnt[i-v]) % mod
			}
		}
	}
	Fprintln(out, cnt[m])
}
