package main

import (
	"math"
	"sort"
)

func minCost(basket1 []int, basket2 []int) int64 {
	cnt := make(map[int]int)
	for i, x := range basket1 {
		cnt[x]++
		cnt[basket2[i]]--
	}
	var arr []int
	mn := math.MaxInt
	for x, c := range cnt {
		if c&1 != 0 {
			return -1
		} else {
			mn = min(mn, x)
			for c = abs(c) >> 1; c != 0; c-- {
				arr = append(arr, x)
			}
		}
	}

	mn <<= 1
	ans := 0
	sort.Ints(arr)
	for _, x := range arr[:len(arr)>>1] {
		ans += min(x, mn)
	}
	return int64(ans)
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
