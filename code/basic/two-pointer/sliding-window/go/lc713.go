package main

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k < 2 {
		return 0
	} else {
		l, product := -1, 1
		for r, x := range nums {
			for product *= x; product >= k; product /= nums[l] {
				l++
			}
			ans += r - l
		}
		return
	}
}
