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
	for _, item := range items { // 枚举每种商品
		v, w, c := item[0], item[1], item[2]
		for ; c > 0; c-- { // 枚举该种商品的每一件，转化为0-1背包
			for j := m; j >= v; j-- { // 枚举背包容量
				dp[j] = max(dp[j], dp[j-v]+w)
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
