package main

import "sort"

func numMovesStonesII(stones []int) []int {
	n := len(stones)
	sort.Ints(stones)
	if stones[n-1]-stones[0] == n-1 { // 所有石子连续
		return []int{0, 0}
	} else {
		le := stones[n-2] - stones[0] - n + 2 // 前n-1颗石子间未占用的位置
		re := stones[n-1] - stones[1] - n + 2 // 后n-1颗石子间未占用的位置
		maxMoves := max(le, re)
		if le == 0 || re == 0 { // 前n-1颗石子连续或后n-1颗石子连续
			return []int{min(2, maxMoves), maxMoves}
		} else {
			minMoves := maxMoves
			for l, r := 0, 0; minMoves != 1 && r < n; r++ {
				for ; stones[r]-stones[l] >= n; l++ {
				}
				minMoves = min(minMoves, n-(r-l+1)) // [stones[r]-n+1, stones[r]]窗口内未占用的位置
			}
			return []int{minMoves, maxMoves}
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

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
