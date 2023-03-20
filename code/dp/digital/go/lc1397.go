package main

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	const mod = 1e9 + 7
	maxMatchLens := calcMaxMatchLens(evil)
	m := len(evil)
	// memo[i][matchLen]记录 f(i, matchLen, false, false) 的结果
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	// 记忆化搜索
	// i: 当前选择从左至右第几位
	// matchLen: 进入第i位之前选择的字符串的后缀与模式串evil的最大匹配长度
	// limitLower: 当前位是否受s1对应位的限制，如果受限则当前位只能选择大于等于s1[i]的字符
	// limitUpper: 当前位是否受s2对应位的限制，如果受限则当前位只能选择小于等于s2[i]的字符
	var f func(i, matchLen int, limitLower, limitUpper bool) int
	f = func(i, matchLen int, limitLower, limitUpper bool) (ans int) {
		if i == n {
			return 1
		} else if !limitLower && !limitUpper && memo[i][matchLen] != -1 {
			// 受限情况下，不能使用记忆的结果，也不能记忆其结果，因为受限情况下的结果是不完整的
			return memo[i][matchLen]
		} else {
			var lower, upper byte = 'a', 'z'
			if limitLower {
				lower = s1[i]
			}
			if limitUpper {
				upper = s2[i]
			}
			for b := lower; b <= upper; b++ {
				ml := matchLen
				for ml > 0 && evil[ml] != b {
					ml = maxMatchLens[ml-1]
				}
				if evil[ml] == b {
					ml++
				}
				if ml < m {
					ll := limitLower && b == lower
					lu := limitUpper && b == upper
					ans = (ans + f(i+1, ml, ll, lu)) % mod
				}
			}
			if !limitLower && !limitUpper {
				memo[i][matchLen] = ans
			}
			return
		}
	}
	return f(0, 0, true, true)
}

// 构造模式串pattern的最大匹配数表
func calcMaxMatchLens(pattern string) []int {
	n := len(pattern)
	maxMatchLens := make([]int, n)
	maxLen := 0
	for i := 1; i < n; i++ {
		for maxLen > 0 && pattern[maxLen] != pattern[i] {
			maxLen = maxMatchLens[maxLen-1]
		}
		if pattern[maxLen] == pattern[i] {
			maxLen++
		}
		maxMatchLens[i] = maxLen
	}
	return maxMatchLens
}
