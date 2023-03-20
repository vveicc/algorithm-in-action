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

	var n, m, p, root int
	Fscan(in, &n, &m)
	og := make([][]int, n+1)
	items := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &items[i][0], &items[i][1], &p)
		if p == -1 {
			root = i
		} else {
			og[p] = append(og[p], i)
		}
	}

	// dp[x][i]表示选择以x为根的子树中的物品，且容量不超过i时能获得的最大价值
	dp := make([][]int, n+1)
	var dfs func(x int)
	dfs = func(x int) {
		v, w := items[x][0], items[x][1]
		dp[x] = make([]int, m+1)
		for i := v; i <= m; i++ {
			dp[x][i] = w // 物品x必须选，初始化为物品x的价值
		}
		for _, y := range og[x] {
			dfs(y)                    // 先计算子树y
			for i := m; i >= v; i-- { // 枚举背包容量
				for j := 0; j <= i-v; j++ { // 枚举留给子树y的背包容量
					dp[x][i] = max(dp[x][i], dp[x][i-j]+dp[y][j])
				}
			}
		}
	}
	dfs(root)

	Fprintln(out, dp[root][m])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
