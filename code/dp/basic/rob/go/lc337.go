package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(o *TreeNode) (with, skip int)
	dfs = func(o *TreeNode) (with int, skip int) {
		if o != nil {
			withL, skipL := dfs(o.Left)
			withR, skipR := dfs(o.Right)
			with = skipL + skipR + o.Val
			skip = max(withL, skipL) + max(withR, skipR)
		}
		return
	}
	with, skip := dfs(root)
	return max(with, skip)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
