package main

import "sort"

func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(limit int) bool {
		cnt, pre := 0, -2
		for i, x := range nums {
			if x <= limit && i-pre != 1 {
				cnt++
				pre = i
			}
		}
		return cnt >= k
	})
}
