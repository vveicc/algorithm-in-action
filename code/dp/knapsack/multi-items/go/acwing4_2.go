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
	for _, item := range items { // 枚举每种物品
		v, w, c := item[0], item[1], item[2]
		for j := m; j >= v; j-- { // 枚举背包容量
			for k := 1; k <= c && k*v <= j; k++ { // 枚举分组内的每一件物品
				dp[j] = max(dp[j], dp[j-k*v]+k*w)
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
