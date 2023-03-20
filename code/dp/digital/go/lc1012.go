package main

import "strconv"

func numDupDigitsAtMostN(n int) int {
	return n - countSpecialNumbers(n)
}

// 统计[1, n]范围内每个数位互不相同的正整数的个数
func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	// memo[i][mask]记录 f(i, mask, false) 的结果
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, 1<<10)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	// 记忆化搜索
	// i: 当前选择从左至右第几位
	// mask: 进入第i位之前的数字选择状态，mask>>d&1 == 1 表示数字 d（0~9）已被选择（有效选择，前导0不算，等同于未选择0）
	// limit: 当前数位是否受n对应数位的限制，例如n=10，则第0个数位只能选择0或1；如果选择1，则第1个数位也受限，只能选择0
	var f func(i, mask int, limit bool) int
	f = func(i, mask int, limit bool) (ans int) {
		if i == m {
			return 1 // 暂且把0（mask = 0）算进去，最终结果再减去1
		} else if !limit && memo[i][mask] != -1 {
			// limit为true时，不能使用记忆的结果，也不能记忆其结果，因为受限情况下的结果是不完整的
			// 例如n=210，进入最右侧数位时，前两位分别选择2、1和1、2对应的mask相同，但是选择2、1是受限状态，最右侧数位只能选择0
			return memo[i][mask]
		} else {
			// 选择0作为当前位
			if mask == 0 { // 当前位选择0只能作为前导0（前导0等同于未使用0，所以mask依然是0）
				ans += f(i+1, mask, false)
			} else if mask&1 == 0 { // 高位已经选了非0的数字，且0未被使用，则当前位可以使用0
				ans += f(i+1, mask|1, limit && s[i] == '0')
			}
			// 选择1~9作为当前位
			upper := 9
			if limit {
				upper = int(s[i] & 15)
			}
			for d := 1; d <= upper; d++ {
				if mask>>d&1 == 0 { // 1~9只要还未被使用，当前位就可以使用
					ans += f(i+1, mask|(1<<d), limit && d == upper)
				}
			}
			if !limit {
				memo[i][mask] = ans
			}
			return
		}
	}
	return f(0, 0, true) - 1
}
