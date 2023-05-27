package main

import "math"

func maxValueAfterReverse(nums []int) int {
	base, d, n := 0, 0, len(nums)
	mn, mx := math.MaxInt, math.MinInt
	for i, x := range nums[1:] {
		mn = min(mn, max(nums[i], x))
		mx = max(mx, min(nums[i], x))
		value := abs(nums[i] - x)
		base += value
		// 边界情况：翻转nums[:i+1]或nums[i+1:]
		d = max(d, max(abs(nums[0]-x), abs(nums[n-1]-nums[i]))-value)
	}
	d = max(d, (mx-mn)*2)
	return base + d
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

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
