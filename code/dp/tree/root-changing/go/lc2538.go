package main

func maxOutput(n int, edges [][]int, price []int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	// 第一遍DFS以节点0为根，计算每个节点向下的最大开销f、次大开销s、最大开销所在子树节点o
	type tuple struct{ f, s, o int }
	values := make([]tuple, n)
	var dfs0 func(x, fa int) int
	dfs0 = func(x, fa int) int {
		f, s, o := 0, 0, -1
		for _, y := range g[x] {
			if y != fa {
				if yf := dfs0(y, x) + price[y]; yf > f {
					f, s, o = yf, f, y
				} else if yf > s {
					s = yf
				}
			}
		}
		values[x] = tuple{f, s, o}
		return f
	}
	ans := dfs0(0, -1)

	// 第二遍DFS通过换根DP计算以每个节点为根的最大开销
	// vf表示节点x往上的最大开销
	var dfs func(x, fa, vf int)
	dfs = func(x, fa, vf int) {
		ans = max(ans, max(vf, values[x].f))
		for _, y := range g[x] {
			if y != fa {
				if y == values[x].o {
					dfs(y, x, max(vf, values[x].s)+price[x])
				} else {
					dfs(y, x, max(vf, values[x].f)+price[x])
				}
			}
		}
	}
	dfs(0, -1, 0)
	return int64(ans)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
