package main

func maximumInvitations(g []int) (ans int) {
	n := len(g)
	// 节点入度
	deg := make([]int, n)
	for _, o := range g {
		deg[o]++
	}

	// 拓扑排序剪掉所有树枝并计算节点最大深度
	var q []int
	for o, d := range deg {
		if d == 0 {
			q = append(q, o)
		}
	}
	maxDepth := make([]int, n)
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		maxDepth[o]++ // 指向o的所有节点已被剪掉
		t := g[o]     // o指向的节点
		maxDepth[t] = max(maxDepth[t], maxDepth[o])
		if deg[t]--; deg[t] == 0 {
			q = append(q, t)
		}
	}

	// 剩余入度不为0的节点都位于环上
	ringSize, ring2Sum := 0, 0
	for o, d := range deg {
		if d != 0 {
			for ringSize = 0; deg[o] != 0; o = g[o] {
				ringSize++
				deg[o] = 0
			}
			if ringSize == 2 {
				ring2Sum += maxDepth[o] + maxDepth[g[o]] + 2
			} else {
				ans = max(ans, ringSize)
			}
		}
	}
	return max(ans, ring2Sum)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
