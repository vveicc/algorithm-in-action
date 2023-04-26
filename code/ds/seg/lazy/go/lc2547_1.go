package main

import "math"

func minCost(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := range nums {
		// state[x]表示x出现的次数，state[x] == 2 表示x已经出现超过一次
		state := make([]int8, n)
		// multi = trimmed(nums[j:i+1]).length
		multi, mn := 0, math.MaxInt
		for j := i; j >= 0; j-- {
			if x := nums[j]; state[x] == 0 {
				state[x]++
			} else if state[x] == 1 {
				state[x]++
				multi += 2
			} else {
				multi++
			}
			mn = min(mn, f[j]+multi)
		}
		f[i+1] = mn + k
	}
	return f[n]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
