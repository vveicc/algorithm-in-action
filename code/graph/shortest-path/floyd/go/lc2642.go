package main

const inf = 1 << 32

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make(Graph, n)
	// 初始化边权
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			if i != j {
				g[i][j] = inf
			}
		}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2]
	}
	// Floyd 算法计算最短路
	for k := range g {
		for x := range g {
			for y := range g {
				g[x][y] = min(g[x][y], g[x][k]+g[k][y])
			}
		}
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	if f, t, d := e[0], e[1], e[2]; d < g[f][t] {
		// Floyd 算法更新最短路
		for x := range g {
			for y := range g {
				g[x][y] = min(g[x][y], g[x][f]+d+g[t][y])
			}
		}
	}
}

func (g Graph) ShortestPath(x int, y int) int {
	if g[x][y] == inf {
		return -1
	} else {
		return g[x][y]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
