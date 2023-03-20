package main

import "strconv"

func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	// 预处理可以选择的数字
	ds := make([]int, len(digits))
	for i, d := range digits {
		ds[i] = int(d[0] & 15)
	}
	// memo[i][selected]记录 f(i, selected, false) 的结果
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = []int{-1, -1}
	}
	// 记忆化搜索
	// i: 当前选择从左至右第几位
	// selected: 进入第i位之前是否已选择digits中的数字，selected==0表示未选择，只有前导0；selected==1表示已选择过
	// limit: 当前数位是否受n对应数位的限制，例如n=10，则第0个数位只能选择0或1；如果选择1，则第1个数位也受限，只能选择0
	var f func(i, selected int, limit bool) int
	f = func(i, selected int, limit bool) (ans int) {
		if i == m {
			return selected // 计算正整数的个数
		} else if !limit && memo[i][selected] != -1 {
			// limit为true时，不能使用记忆的结果，也不能记忆其结果，因为受限情况下的结果是不完整的
			return memo[i][selected]
		} else {
			if selected == 0 {
				ans += f(i+1, selected, false) // 当前数位继续不选择
			}
			upper := 9
			if limit {
				upper = int(s[i] & 15)
			}
			for _, d := range ds {
				if d > upper {
					break
				} else {
					ans += f(i+1, 1, limit && d == upper) // 当前数位选择数字d
				}
			}
			if !limit {
				memo[i][selected] = ans
			}
			return
		}
	}
	return f(0, 0, true)
}
