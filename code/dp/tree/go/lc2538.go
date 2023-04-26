package main

func maxOutput(n int, edges [][]int, price []int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	ans := 0
	// 返回子树带叶子节点的最大路径和、不带叶子节点的最大路径和
	var dfs func(x, fa int) (int, int)
	dfs = func(x, fa int) (int, int) {
		p := price[x]
		maxS1, maxS2 := p, 0
		for _, y := range g[x] {
			if y != fa {
				s1, s2 := dfs(y, x)
				ans = max(ans, max(maxS1+s2, maxS2+s1))
				maxS1 = max(maxS1, s1+p)
				maxS2 = max(maxS2, s2+p)
			}
		}
		return maxS1, maxS2
	}
	dfs(0, -1)
	return int64(ans)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
