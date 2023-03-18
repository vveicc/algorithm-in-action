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
	// groups[i][j]表示从第i个dq中取出j+1个数的最大和
	groups := make([][]int, n)
	for i := 0; i < n; i++ {
		var k, x int
		Fscan(in, &k)
		s := make([]int, k+1) // 前缀和
		for j := 0; j < k; j++ {
			Fscan(in, &x)
			s[j+1] = s[j] + x
		}
		groups[i] = make([]int, k)
		groups[i][k-1] = s[k]
		for j := 1; j < k; j++ {
			for l := 0; l <= j; l++ { // 在dq头部取l个数，尾部取j-l个数
				groups[i][j-1] = max(groups[i][j-1], s[l]+s[k]-s[k-(j-l)])
			}
		}
	}

	t := 0
	dp := make([]int, m+1)
	for _, group := range groups { // 枚举分组
		k := len(group)
		t += k
		for i := min(t, m); i > 0; i-- { // 枚举背包容量
			for j := 0; j < k && j < i; j++ { // 枚举分组内的每一件物品
				// 使用j+1的容量容纳group[j]
				dp[i] = max(dp[i], dp[i-1-j]+group[j])
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
