package main

func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)
	ug := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		ug[x] = append(ug[x], y)
		ug[y] = append(ug[y], x)
	}

	const inf = int(-1e8)

	// 第一遍DFS以节点0为根，计算每个节点与其子树中的最远金币距离f、次远金币距离s、最远金币所在子树节点o
	// 同时计算出从节点0出发，收集所有金币需要经过的单程边数（不返回出发节点）
	type tuple struct{ f, s, o int }
	dis := make([]tuple, n)
	var dfs0 func(x, fa int) (int, e int)
	dfs0 = func(x, fa int) (int, e int) {
		// 距离初始化为inf，方便处理子树中没有金币的情况
		f, s, o := inf, inf, -1
		for _, y := range ug[x] {
			if y != fa {
				yf, ye := dfs0(y, x)
				e += ye
				if yf++; yf > 2 {
					e++ // 需要经过x-y这条边
				}
				if yf > f {
					f, s, o = yf, f, y
				} else if yf > s {
					s = yf
				}
			}
		}
		dis[x] = tuple{f, s, o}
		if f < 0 && coins[x] == 1 {
			f = 0
		}
		return f, e
	}
	_, ans := dfs0(0, -1)

	// 第二遍DFS通过换根DP计算从每个节点出发收集所有金币需要经过的单程边数（不返回出发节点）
	// d表示节点x往上的最远金币距离
	var dfs func(x, fa, d, e int)
	dfs = func(x, fa, d, e int) {
		if ans = min(ans, e); ans == 0 {
			return // 剪枝
		}
		if d < 0 && coins[x] == 1 {
			d = 0
		}
		for _, y := range ug[x] {
			if y != fa && dis[y].f >= 2 { // 从x节点出发需要经过x-y这条边
				var yd int
				if y == dis[x].o {
					yd = max(d, dis[x].s) + 1
				} else {
					yd = max(d, dis[x].f) + 1
				}
				if yd <= 2 {
					dfs(y, x, yd, e-1) // 从y节点出发不需要经过x-y这条边
				}
			}
		}
	}
	dfs(0, -1, inf, ans)
	return ans << 1
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
