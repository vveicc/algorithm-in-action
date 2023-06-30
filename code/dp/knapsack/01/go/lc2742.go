package main

import "math"

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i, c := range cost { // 枚举物品
		t := time[i] + 1         // 物品体积
		for j := n; j > 0; j-- { // 枚举背包容量
			dp[j] = min(dp[j], dp[max(j-t, 0)]+c)
		}
	}
	return dp[n]
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
