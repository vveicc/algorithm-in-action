package main

import "math"

func minCost(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := range nums {
		// state[x]表示x出现的次数，state[x] == 2 表示x已经出现超过一次
		state := make([]int8, n)
		unique, mn := 0, math.MaxInt
		for j := i; j >= 0; j-- {
			if x := nums[j]; state[x] == 0 { // 首次出现
				state[x] = 1
				unique++
			} else if state[x] == 1 { // 不再唯一
				state[x] = 2
				unique--
			}
			mn = min(mn, f[j]-unique)
		}
		f[i+1] = mn + k
	}
	return f[n] + n
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
