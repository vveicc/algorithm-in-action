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
			mask, d := s, 0
			var dfs func(x, fa int) int
			dfs = func(x, fa int) int {
				mask ^= 1 << x
				var f, s int // 最大子树深度和次大子树深度
				for _, y := range ug[x] {
					if y != fa && mask>>y&1 == 1 {
						if yf := dfs(y, x) + 1; yf > f {
							f, s = yf, f
						} else if yf > s {
							s = yf
						}
					}
				}
				d = max(d, f+s)
				return f // 返回最大子树深度
			}
			start := bits.Len16(uint16(s)) - 1
			if dfs(start, -1); mask == 0 { // 判断子树s是否有效
				ans[d-1]++
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
