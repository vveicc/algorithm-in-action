package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) (ans int) {
	var arr [][]int
	var dfs func(o *TreeNode, d int)
	dfs = func(o *TreeNode, d int) {
		if o != nil {
			if d == len(arr) {
				arr = append(arr, []int{o.Val})
			} else {
				arr[d] = append(arr[d], o.Val)
			}
			dfs(o.Left, d+1)
			dfs(o.Right, d+1)
		}
	}
	dfs(root, 0)
	for _, values := range arr[1:] {
		if n := len(values); n > 1 {
			// 离散化
			id := make([]int, n)
			for i := range id {
				id[i] = i
			}
			sort.Slice(id, func(i, j int) bool { return values[id[i]] < values[id[j]] })

			// 置换环
			ans += n
			vis := make([]bool, n)
			for _, i := range id {
				if !vis[i] {
					for ans--; !vis[i]; i = id[i] {
						vis[i] = true
					}
				}
			}
		}
	}
	return
}
