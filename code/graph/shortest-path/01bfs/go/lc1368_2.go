package main

import "container/heap"

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	mn := m * n
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = mn
		}
	}
	h := hp{{}}
	dis[0][0] = 0
	for h.Len() > 0 {
		f := heap.Pop(&h).(pair)
		if x, y := f.x, f.y; f.dis <= dis[x][y] {
			for i, dir := range dirs {
				tx, ty := x+dir[0], y+dir[1]
				if 0 <= tx && tx < m && 0 <= ty && ty < n {
					cost := 0
					if i+1 != grid[x][y] {
						cost = 1
					}
					if d := f.dis + cost; d < dis[tx][ty] {
						dis[tx][ty] = d
						heap.Push(&h, pair{tx, ty, d})
					}
				}
			}
		}
	}
	return dis[m-1][n-1]
}

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type pair struct{ x, y, dis int } // 单元格(x, y)及其与起点的距离
type hp []pair                    // 堆（优先队列）

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
