package main

import (
	"sort"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
	n := len(robot)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1e12
	}
	acc := 0
	for _, f := range factory {
		acc += f[1]
		for j := min(n, acc); j > 0; j-- {
			// 在当前工厂可以维修0~cnt个机器人
			cnt := min(j, f[1])
			// 前j个机器人中的后k个进入当前工厂
			for k, dis := 1, 0; k <= cnt; k++ {
				dis += abs(f[0] - robot[j-k])
				dp[j] = min(dp[j], dp[j-k]+dis)
			}
		}
	}
	return int64(dp[n])
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
