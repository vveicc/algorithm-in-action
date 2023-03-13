package main

import (
	"container/heap"
	"sort"
)

func findCrossingTime(n int, k int, time [][]int) (curr int) {
	// 按照优先级从低到高排序，排序完成后，索引越大效率越低，过桥优先级越高
	sort.SliceStable(time, func(i, j int) bool {
		x, y := time[i], time[j]
		return x[0]+x[2] < y[0]+y[2]
	})
	// 左右两岸正在放箱/搬箱的工人
	workL, workR := working{}, working{}
	// 左右两岸正在等待过桥的工人
	waitL, waitR := make(waiting, k), waiting{}
	for i := range waitL {
		waitL[i].i = k - 1 - i // 索引越大效率越低，过桥优先级越高
	}
	for n > 0 {
		// 左岸完成放箱，进入左岸等待过桥队列
		for workL.Len() > 0 && workL[0].t <= curr {
			heap.Push(&waitL, heap.Pop(&workL))
		}
		// 右岸完成搬箱，进入右岸等待过桥队列
		for workR.Len() > 0 && workR[0].t <= curr {
			heap.Push(&waitR, heap.Pop(&workR))
		}
		if waitR.Len() > 0 && waitR[0].t <= curr { // 右岸等待过桥的工人过桥
			w := heap.Pop(&waitR).(worker)
			curr += time[w.i][2]                                // 时间来到w过桥结束来到左岸的时刻
			heap.Push(&workL, worker{w.i, curr + time[w.i][3]}) // w进入放箱队列，记录放箱完成时间
		} else if waitL.Len() > 0 && waitL[0].t <= curr { // 左岸等待过桥的工人过桥
			w := heap.Pop(&waitL).(worker)
			curr += time[w.i][0]                                // 时间来到w过桥结束来到右岸的时刻
			heap.Push(&workR, worker{w.i, curr + time[w.i][1]}) // w进入搬箱队列，记录搬箱完成时间
			n--                                                 // w将会完成一个箱子的搬运
		} else {
			// 此时所有工人都在进行放箱/搬箱的工作，没有等待过桥的工人
			// 我们让无聊的时间来到所有工人中最早完成当前工作准备过桥的时刻
			if workL.Len() == 0 {
				curr = workR[0].t
			} else if workR.Len() == 0 {
				curr = workL[0].t
			} else {
				curr = min(workL[0].t, workR[0].t)
			}
		}
	}
	// 此时搬运最后一个箱子的工人已经来到右岸
	// 等待所有右岸的工人完成搬箱并过桥前往左岸即可，由于所有右岸的工人都要前往左岸，所以过桥的优先级顺序已不再重要
	for workR.Len() > 0 {
		w := heap.Pop(&workR).(worker)
		curr = max(curr, w.t) + time[w.i][2]
	}
	return
}

type worker struct{ i, t int } // 工人编号及完成时间

type working []worker

func (h working) Len() int            { return len(h) }
func (h working) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h working) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *working) Push(x interface{}) { *h = append(*h, x.(worker)) }
func (h *working) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type waiting []worker

func (h waiting) Len() int            { return len(h) }
func (h waiting) Less(i, j int) bool  { return h[i].i > h[j].i }
func (h waiting) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *waiting) Push(x interface{}) { *h = append(*h, x.(worker)) }
func (h *waiting) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
