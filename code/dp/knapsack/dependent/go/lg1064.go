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
	n /= 10
	items := make([][3]int, m+1)
	og := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &items[i][0], &items[i][1], &items[i][2])
		items[i][0] /= 10
		items[i][1] *= items[i][0]
		if q := items[i][2]; q != 0 {
			og[q] = append(og[q], i)
		}
	}
	dp := make([]int, n+1)
	for x := 1; x <= m; x++ { // 枚举主件与其附件组成的分组
		if items[x][2] == 0 {
			vx, wx := items[x][0], items[x][1]
			for i := n; i >= vx; i-- { // 枚举背包容量
				for s := 1<<len(og[x]) - 1; s >= 0; s-- { // 枚举分组内的物品（附件选择状态）
					v, w := vx, wx // 选择主件x
					for j, y := range og[x] {
						if s>>j&1 == 1 { // 选择状态s表示的附件
							v += items[y][0]
							w += items[y][1]
						}
					}
					if i >= v {
						dp[i] = max(dp[i], dp[i-v]+w)
					}
				}
			}
		}
	}

	Fprintln(out, dp[n]*10)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
