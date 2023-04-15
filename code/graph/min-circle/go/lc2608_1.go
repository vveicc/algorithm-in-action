package main

import "math"

func findShortestCycle(n int, edges [][]int) int {
	ug := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		ug[u] = append(ug[u], v)
		ug[v] = append(ug[v], u)
	}
	ans := math.MaxInt
	dis := make([]int, n)
	for _, e := range edges {
		f, t := e[0], e[1]
		// 计算从节点f出发到达节点t，且不经过边f-t的最短路
		// 由于边权均为1，通过BFS就可以计算最短路
		for i := range dis {
			dis[i] = -1
		}
		dis[f] = 0
		q := []int{f}
		for len(q) != 0 {
			x := q[0]
			q = q[1:]
			for _, y := range ug[x] {
				if dis[y] == -1 && (x != f || y != t) {
					dis[y] = dis[x] + 1
					q = append(q, y)
				}
			}
		}
		if dis[t] != -1 {
			ans = min(ans, dis[t]+1)
		}
	}
	if ans == math.MaxInt {
		return -1
	} else {
		return ans
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
