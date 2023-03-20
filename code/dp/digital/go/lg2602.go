package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b int
	Fscan(in, &a, &b)
	low, high := strconv.Itoa(a), strconv.Itoa(b)
	m, n := len(low), len(high)
	if m < n {
		low = strings.Repeat("0", n-m) + low
	}
	for d := byte('0'); d <= '9'; d++ {
		Fprint(out, countDigit(d, n, low, high), " ")
	}
}

func countDigit(d byte, n int, low, high string) int {
	// memo[i][cnt]记录 f(i, cnt, false, false, false) 的结果
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	// 记忆化搜索
	// i: 当前选择从左至右第几位
	// cnt: 进入第i位之前数字d出现的次数
	// zero: 进入第i位之前是否未选择非0数字，如果未选择则当前位可以选择0作为前缀0
	// limitLower: 当前位是否受low对应位的限制，如果受限则当前位只能选择大于等于low[i]的字符
	// limitUpper: 当前位是否受high对应位的限制，如果受限则当前位只能选择小于等于high[i]的字符
	var f func(i, cnt int, zero, limitLower, limitUpper bool) int
	f = func(i, cnt int, zero, limitLower, limitUpper bool) (ans int) {
		if i == n {
			return cnt
		} else if !zero && !limitLower && !limitUpper && memo[i][cnt] != -1 {
			// 受限情况下，不能使用记忆的结果，也不能记忆其结果，因为受限情况下的结果是不完整的
			return memo[i][cnt]
		} else {
			var lower, upper byte = '0', '9'
			if limitLower {
				lower = low[i]
			}
			if limitUpper {
				upper = high[i]
			}
			for b := lower; b <= upper; b++ {
				zo := zero && b == '0'
				ll := limitLower && b == lower
				lu := limitUpper && b == upper
				if !zo && b == d {
					ans += f(i+1, cnt+1, zo, ll, lu)
				} else {
					ans += f(i+1, cnt, zo, ll, lu)
				}
			}
			if !zero && !limitLower && !limitUpper {
				memo[i][cnt] = ans
			}
			return ans
		}
	}
	return f(0, 0, true, true, true)
}
