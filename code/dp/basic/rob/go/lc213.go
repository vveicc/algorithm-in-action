package main

func rob(nums []int) int {
	if n := len(nums); n == 1 {
		return nums[0]
	} else {
		f := func(arr []int) int {
			with, skip := 0, 0
			for _, x := range arr {
				with, skip = skip+x, max(with, skip)
			}
			return max(with, skip)
		}
		return max(f(nums[1:]), f(nums[:n-1]))
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
