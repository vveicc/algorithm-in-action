package main

import "strconv"

func findIntegers(n int) int {
	// 转为二进制字符串
	s := strconv.FormatInt(int64(n), 2)
	m := len(s)
	// memo[i][prev]记录 f(i, prev, false) 的结果
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = []int{-1, -1}
	}
	// 记忆化搜索
	// i: 当前选择从左至右第几位
	// prev: 前一个2进制数位选择的数字
	// limit: 当前数位是否受n对应数位的限制，例如n=10，则第0个数位只能选择0或1；如果选择1，则第1个数位也受限，只能选择0
	var f func(i, prev int, limit bool) int
	f = func(i, prev int, limit bool) (ans int) {
		if i == m {
			return 1
		} else if !limit && memo[i][prev] != -1 {
			return memo[i][prev]
		} else {
			// 选择0
			ans += f(i+1, 0, limit && s[i] == '0')
			// 选择1
			if prev == 0 && (!limit || s[i] == '1') {
				ans += f(i+1, 1, limit && s[i] == '1')
			}
			if !limit {
				memo[i][prev] = ans
			}
			return ans
		}
	}
	return f(0, 0, true)
}
