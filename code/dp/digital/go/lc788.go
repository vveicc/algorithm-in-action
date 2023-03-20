package main

import "strconv"

// score[d]==1表示至少需要包含一个数字d，score[d]==-1表示不能包含数字d
var score = []int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}

func rotatedDigits1(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	// memo[i][diff]记录 f(i, diff, false) 的结果
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = []int{-1, -1}
	}
	// 记忆化搜索
	// i: 当前选择从左至右第几位
	// diff: 进入第i位之前是否已选择过2、5、6、9，diff==0表示未选择过，diff==1表示选择过
	// limit: 当前数位是否受n对应数位的限制，例如n=10，则第0个数位只能选择0或1；如果选择1，则第1个数位也受限，只能选择0
	var f func(i, diff int, limit bool) int
	f = func(i, diff int, limit bool) (ans int) {
		if i == m {
			return diff
		} else if !limit && memo[i][diff] != -1 {
			// limit为true时，不能使用记忆的结果，也不能记忆其结果，因为受限情况下的结果是不完整的
			return memo[i][diff]
		} else {
			upper := 9
			if limit {
				upper = int(s[i] & 15)
			}
			for d := 0; d <= upper; d++ {
				if score[d] != -1 {
					ans += f(i+1, diff|score[d], limit && d == upper)
				}
			}
			if !limit {
				memo[i][diff] = ans
			}
			return ans
		}
	}
	return f(0, 0, true)
}
