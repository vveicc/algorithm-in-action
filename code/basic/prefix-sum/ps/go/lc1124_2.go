package main

func longestWPI(hours []int) int {
	for i, x := range hours {
		if x > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}
	return longestSubarray(hours, 1)
}

// 泛化：和至少为 K 的最长子数组
func longestSubarray(nums []int, k int) (ans int) {
	k-- // 和至少为k -> 和大于k
	n := len(nums)
	s := make([]int, n+1)   // 前缀和
	st, tail := []int{}, -1 // 前缀和单调递减索引栈
	for i, x := range nums {
		i++
		if s[i] = s[i-1] + x; s[i] > k {
			ans = i
		}
		if tail == -1 || s[i] < s[st[tail]] {
			st = append(st, i)
			tail++
		}
	}
	for i := n; i > ans && tail != -1; i-- {
		for ; tail != -1 && s[i]-s[st[tail]] > k; tail-- {
			ans = max(ans, i-st[tail])
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
