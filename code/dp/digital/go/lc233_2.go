package main

import "strconv"

func countDigitOne(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	// memo[i][cnt]记录 f(i, cnt, false) 的结果
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	// 记忆化搜索
	// i: 当前选择从左至右第几位
	// cnt: 进入第i位之前数字1出现的次数
	// limit: 当前数位是否受n对应数位的限制，例如n=10，则第0个数位只能选择0或1；如果选择1，则第1个数位也受限，只能选择0
	var f func(i, cnt int, limit bool) int
	f = func(i, cnt int, limit bool) (ans int) {
		if i == m {
			return cnt
		} else if !limit && memo[i][cnt] != -1 {
			// limit为true时，不能使用记忆的结果，也不能记忆其结果，因为受限情况下的结果是不完整的
			return memo[i][cnt]
		} else {
			upper := 9
			if limit {
				upper = int(s[i] & 15)
			}
			for d := 0; d <= upper; d++ {
				if d == 1 {
					ans += f(i+1, cnt+1, limit && d == upper)
				} else {
					ans += f(i+1, cnt, limit && d == upper)
				}
			}
			if !limit {
				memo[i][cnt] = ans
			}
			return
		}
	}
	return f(0, 0, true)
}
