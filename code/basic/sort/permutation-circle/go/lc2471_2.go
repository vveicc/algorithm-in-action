package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) (ans int) {
	var tmp []*TreeNode
	q := []*TreeNode{root}
	for n := len(q); n > 0; n = len(q) {
		tmp, q = q, nil
		for _, o := range tmp {
			if o.Left != nil {
				q = append(q, o.Left)
			}
			if o.Right != nil {
				q = append(q, o.Right)
			}
		}
		if n != 1 {
			// 离散化
			id := make([]int, n)
			arr := make([]int, n)
			for i, o := range tmp {
				id[i] = i
				arr[i] = o.Val
			}
			sort.Slice(id, func(i, j int) bool { return arr[id[i]] < arr[id[j]] })

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
