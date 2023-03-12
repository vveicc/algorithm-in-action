package main

import "strings"

func minimumRecolors(blocks string, k int) int {
	ans := strings.Count(blocks[:k], "W")
	countW := ans // 记录长为k的滑动窗口内白色块的数量
	for i, n := k, len(blocks); i < n; i++ {
		// 'W'&1 = 1; 'B'&1 = 0
		countW += int(blocks[i]&1) - int(blocks[i-k]&1)
		ans = min(ans, countW)
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
