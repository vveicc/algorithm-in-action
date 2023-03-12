package main

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	ug := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		ug[u] = append(ug[u], v)
		ug[v] = append(ug[v], u)
	}

	// dis[i][j]表示节点i和j的距离
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, n)
		// 以节点i为根计算每个节点的深度，即与节点i的距离
		var dfs func(x, fa int)
		dfs = func(x, fa int) {
			for _, y := range ug[x] {
				if y != fa {
					dis[i][y] = dis[i][x] + 1 // 自顶向下
					dfs(y, x)
				}
			}
		}
		dfs(i, -1)
	}

	ans := make([]int, n-1)
	for i, di := range dis {
		for j := i + 1; j < n; j++ {
			d := di[j]
			dj := dis[j]
			var dfs func(x, fa int) int
			dfs = func(x, fa int) int {
				// 选择x
				cnt := 1
				for _, y := range ug[x] {
					if y != fa &&
						(di[y] < d || (di[y] == d && y > j)) &&
						(dj[y] < d || (dj[y] == d && y > i)) { // 可以选择节点y
						cnt *= dfs(y, x) // 每颗子树相互独立，使用乘法原理
					}
				}
				if di[x]+dj[x] > d { // 节点x非必选
					cnt++ // 不选择x
				}
				return cnt
			}
			ans[d-1] += dfs(i, -1)
		}
	}
	return ans
}
