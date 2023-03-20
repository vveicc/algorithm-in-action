package main

import "strconv"

func findIntegers(n int) int {
	// 转为4进制字符串
	s := strconv.FormatInt(int64(n), 4)
	m := len(s)
	// memo[i][prev]记录 f(i, prev, false) 的结果
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, 4)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	// 记忆化搜索
	// i: 当前选择从左至右第几位
	// prev: 前一个4进制数位选择的数字
	// limit: 当前数位是否受n对应数位的限制，例如n=10，则第0个数位只能选择0或1；如果选择1，则第1个数位也受限，只能选择0
	var f func(i, prev int, limit bool) int
	f = func(i, prev int, limit bool) (ans int) {
		if i == m {
			return 1
		} else if !limit && memo[i][prev] != -1 {
			// limit为true时，不能使用记忆的结果，也不能记忆其结果，因为受限情况下的结果是不完整的
			return memo[i][prev]
		} else {
			upper := 3
			if limit {
				upper = int(s[i] & 15)
			}
			for d := 0; d <= upper && d < 3; d++ {
				if d < 2 || prev != 1 { // 当前4进制位可以是0、1、2，且前一位为1时，当前位不能是2
					ans += f(i+1, d, limit && d == upper)
				}
			}
			if !limit {
				memo[i][prev] = ans
			}
			return
		}
	}
	return f(0, 0, true)
}
