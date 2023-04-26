package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt64
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		} else {
			maxGainL := max(dfs(node.Left), 0)
			maxGainR := max(dfs(node.Right), 0)
			ans = max(ans, node.Val+maxGainL+maxGainR)
			return node.Val + max(maxGainL, maxGainR)
		}
	}
	dfs(root)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
