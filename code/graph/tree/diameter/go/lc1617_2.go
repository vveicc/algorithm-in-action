package main

import "math/bits"

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	ug := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		ug[u] = append(ug[u], v)
		ug[v] = append(ug[v], u)
	}
	ans := make([]int, n-1)
	// 状态压缩枚举城市子集
	for s := 1<<n - 1; s != 0; s-- {
		if bits.OnesCount16(uint16(s)) != 1 { // 至少需要两个节点
			max, o, vis := 0, -1, 0
			var dfs func(x, fa, depth int)
			dfs = func(x, fa, depth int) {
				vis |= 1 << x
				if depth > max {
					max, o = depth, x
				}
				for _, y := range ug[x] {
					if y != fa && s>>y&1 == 1 {
						dfs(y, x, depth+1)
					}
				}
			}
			// 第一遍DFS找到与start节点距离最远的节点
			start := bits.Len16(uint16(s)) - 1
			if dfs(start, -1, 0); vis == s { // 判断子树s是否有效
				max = 0
				dfs(o, -1, 0) // 第二遍DFS找到与o节点距离最远的节点，两节点距离即为子树的直径
				ans[max-1]++
			}
		}
	}
	return ans
}
