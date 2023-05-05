package main

import "container/heap"

func minimumCost(start []int, target []int, specialRoads [][]int) int {
	n := len(specialRoads)
	sx, sy := start[0], start[1]
	tx, ty := target[0], target[1]
	dis := make([]int, n)
	vis := make([]bool, n+1)
	for i, r := range specialRoads {
		x2, y2 := r[2], r[3]
		dis[i] = x2 - sx + y2 - sy
	}
	ans := tx - sx + ty - sy
	h := hp{{n<<shift2 | sx<<shift | sy, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		if i, x, y := p.ixy>>shift2, p.ixy>>shift&mask, p.ixy&mask; !vis[i] {
			vis[i] = true
			ans = min(ans, p.dis+tx-x+ty-y)
			for j, r := range specialRoads {
				if !vis[j] {
					x1, y1, x2, y2, dj := r[0], r[1], r[2], r[3], r[4]
					if d := p.dis + abs(x-x1) + abs(y-y1) + dj; d < dis[j] {
						dis[j] = d
						heap.Push(&h, pair{j<<shift2 | x2<<shift | y2, d})
					}
				}
			}
		}
	}
	return ans
}

const shift, shift2, mask = 20, 40, 1<<20 - 1

type pair struct{ ixy, dis int }
type hp []pair

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

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
