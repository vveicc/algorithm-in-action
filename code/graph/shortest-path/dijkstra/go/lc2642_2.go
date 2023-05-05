package main

import "container/heap"

const inf = 1 << 32

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make([][]int, n)
	for _, e := range edges {
		f, t, d := e[0], e[1], e[2]
		g[f] = append(g[f], d<<shift|t)
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	f, t, d := e[0], e[1], e[2]
	g[f] = append(g[f], d<<shift|t)
}

func (g Graph) ShortestPath(start int, end int) int {
	n := len(g)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		if f := p.o; f == end {
			return dis[f]
		} else if p.dis == dis[f] {
			for _, dt := range g[f] {
				t := dt & mask
				if d := p.dis + dt>>shift; d < dis[t] {
					dis[t] = d
					heap.Push(&h, pair{t, d})
				}
			}
		}
	}
	return -1
}

const shift, mask = 8, 1<<8 - 1

type pair struct{ o, dis int }
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

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
