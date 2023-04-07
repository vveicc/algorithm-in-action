package main

func rob(nums []int) int {
	with, skip := 0, 0
	for _, x := range nums {
		with, skip = skip+x, max(with, skip)
	}
	return max(with, skip)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
