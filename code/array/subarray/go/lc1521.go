package main

import "math"

func closestToTarget(arr []int, target int) int {
	ans := math.MaxInt
	var a []int // 记录每个连续段的按位与值
	for _, x := range arr {
		// 计算每个连续段的按位与值
		for j := range a {
			a[j] &= x
		}
		// 增加新的连续段
		a = append(a, x)
		// 原地去重
		j := 0
		ans = min(ans, abs(a[0]-target))
		for _, y := range a[1:] {
			if y != a[j] {
				j++
				a[j] = y
				ans = min(ans, abs(y-target))
			}
		}
		a = a[:j+1]
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
