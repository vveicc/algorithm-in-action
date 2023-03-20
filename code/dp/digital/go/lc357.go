package main

func countNumbersWithUniqueDigits(n int) int {
	// memo[i][mask]记录 f(i, mask) 的结果
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<10)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	// 记忆化搜索
	// i: 当前选择从左至右第几位
	// mask: 进入第i位之前的数字选择状态，mask>>d&1 == 1 表示数字 d（0~9）已被选择（有效选择，前导0不算，等同于未选择0）
	var f func(i, mask int) int
	f = func(i, mask int) (ans int) {
		if i == n {
			return 1
		} else if memo[i][mask] != -1 {
			return memo[i][mask]
		} else {
			// 选择0作为当前位
			if mask == 0 { // 当前位选择0只能作为前导0（前导0等同于未使用0，所以mask依然是0）
				ans += f(i+1, mask)
			} else if mask&1 == 0 { // 高位已经选了非0的数字，且0未被使用，则当前位可以使用0
				ans += f(i+1, mask|1)
			}
			// 选择1~9作为当前位
			for d := 1; d <= 9; d++ {
				if mask>>d&1 == 0 {
					ans += f(i+1, mask|(1<<d))
				}
			}
			memo[i][mask] = ans
			return ans
		}
	}
	return f(0, 0)
}
