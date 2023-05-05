package main

const inf = 1 << 32

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make(Graph, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2]
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]][e[1]] = e[2]
}

func (g Graph) ShortestPath(start int, end int) int {
	n := len(g)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	vis := make([]bool, n)
	for {
		x := -1
		for i, b := range vis {
			if !b && (x == -1 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x == -1 || dis[x] == inf {
			return -1
		}
		if x == end {
			return dis[x]
		}
		vis[x] = true
		for y, d := range g[x] {
			dis[y] = min(dis[y], dis[x]+d)
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
