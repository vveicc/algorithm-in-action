package main

import "container/heap"

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	mn := m * n
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = mn
		}
	}
	dis[0][0] = 0

	h := hp{{}}
	for h.Len() != 0 {
		f := heap.Pop(&h).(tuple)
		if x, y := f.x, f.y; f.dis <= dis[x][y] {
			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]
				if 0 <= nx && nx < m && 0 <= ny && ny < n {
					if d := f.dis + grid[nx][ny]; d < dis[nx][ny] {
						dis[nx][ny] = d
						heap.Push(&h, tuple{nx, ny, d})
					}
				}
			}
		}
	}
	return dis[m-1][n-1]
}

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type tuple struct{ x, y, dis int } // 单元格(x, y)及其与起点的距离
type hp []tuple                    // 堆（优先队列）

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(tuple)) }
func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
