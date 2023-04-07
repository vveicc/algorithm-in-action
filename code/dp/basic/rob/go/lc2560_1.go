package main

import "sort"

func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(limit int) bool {
		f0, f1 := 0, 0
		for _, x := range nums {
			if x > limit {
				f0 = f1
			} else {
				f0, f1 = f1, f0+1
			}
		}
		return f1 >= k
	})
}
