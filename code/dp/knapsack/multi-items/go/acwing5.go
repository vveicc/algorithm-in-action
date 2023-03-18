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
	items := make([][3]int, n)
	for i := range items {
		Fscan(in, &items[i][0], &items[i][1], &items[i][2])
	}

	dp := make([]int, m+1)
	for _, item := range items {
		v, w, c := item[0], item[1], item[2]
		if c*v >= m { // 转化为完全背包
			for j := v; j <= m; j++ {
				dp[j] = max(dp[j], dp[j-v]+w)
			}
		} else { // 不能转化为完全背包，使用二进制分组优化
			// 每个分组看作独立的一件物品，转化为0-1背包，组合分组可以组合出[1, c]范围内的任一件数
			for k := 1; k <= c; k <<= 1 {
				c -= k
				kv, kw := k*v, k*w
				for j := m; j >= kv; j-- {
					dp[j] = max(dp[j], dp[j-kv]+kw)
				}
			}
			if c != 0 {
				cv, cw := c*v, c*w
				for j := m; j >= cv; j-- {
					dp[j] = max(dp[j], dp[j-cv]+cw)
				}
			}
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
