package main

import "sort"

func minDistance(houses []int, k int) int {
	n := len(houses)
	sort.Ints(houses)

	// medsum[i][j]表示有序点集houses[i:j+1]到中位数的最小距离总和
	medsum := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		medsum[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			medsum[i][j] = medsum[i+1][j-1] + houses[j] - houses[i]
		}
	}

	// dp[i][j]表示给编号为[0..i]的前i+1个房子，安排编号为[0..j]的j+1个邮筒的最小距离总和
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
		// 安排编号为0的邮筒
		dp[i][0] = medsum[0][i]
		// 安排编号为j的邮筒
		for j := 1; j < i && j < k; j++ {
			dp[i][j] = 1e6
			// 枚举与编号为j的邮筒距离最近的房子的起始编号i0
			for i0 := j; i0 <= i; i0++ {
				dp[i][j] = min(dp[i][j], dp[i0-1][j-1]+medsum[i0][i])
			}
		}
	}
	return dp[n-1][k-1]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
