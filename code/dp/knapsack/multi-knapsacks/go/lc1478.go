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

	// dp[i][j]表示给houses[:j+1]安排编号为[0..i]的i+1个邮筒的最小距离总和
	dp := make([]int, n)
	// 安排编号为0的1个邮筒
	for j := range dp {
		dp[j] = medsum[0][j]
	}
	// 安排编号为i的第i+1个邮筒
	for i := 1; i < k; i++ {
		// 给houses[:j+1]安排编号为[0..i]的i+1个邮筒
		// j == i表示给houses[:i+1]每个房子安排一个邮筒
		for j := n - 1; j >= i; j-- {
			// 将houses[j0:j+1]安排给编号为i的第i+1个邮筒
			for j0 := i; j0 <= j; j0++ {
				dp[j] = min(dp[j], dp[j0-1]+medsum[j0][j])
			}
		}
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
