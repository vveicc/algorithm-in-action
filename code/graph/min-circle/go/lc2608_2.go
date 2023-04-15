package main

import "math"

const shift, mask = 10, 1<<10 - 1

func findShortestCycle(n int, edges [][]int) int {
	ug := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		ug[u] = append(ug[u], v)
		ug[v] = append(ug[v], u)
	}
	ans := math.MaxInt
	dis := make([]int, n)
	for o := 0; o < n; o++ {
		for i := range dis {
			dis[i] = -1
		}
		dis[o] = 0
		q := []int{o<<shift | o}
		for len(q) > 0 {
			f, x := q[0]>>shift, q[0]&mask
			q = q[1:]
			for _, y := range ug[x] {
				if dis[y] == -1 { // 第一次访问节点y，计算与节点o的最短距离，并入队
					dis[y] = dis[x] + 1
					q = append(q, x<<shift|y)
				} else if y != f { // 第二次访问节点y，通过判断排除y-x-y这种不成环的访问情形
					ans = min(ans, dis[x]+dis[y]+1)
				}
			}
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
