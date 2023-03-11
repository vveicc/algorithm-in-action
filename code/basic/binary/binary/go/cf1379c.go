package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, m, a, b int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		abs := make([][2]int, m)
		for i := 0; i < m; i++ {
			Fscan(in, &a, &b)
			abs[i] = [2]int{a, b}
		}

		// 按照a降序排序
		sort.Slice(abs, func(i, j int) bool { return abs[i][0] > abs[j][0] })
		// 前缀和
		s := make([]int64, m+1)
		for i, ab := range abs {
			s[i+1] = s[i] + int64(ab[0])
		}
		var ans int64
		for _, ab := range abs {
			a, b = ab[0], ab[1]
			// 前i个物品的a大于当前物品的b，如果当前物品取多个（多于1个），则前i个物品需要每个取1个
			i := sort.Search(m, func(i int) bool { return abs[i][0] <= b })
			if i >= n {
				ans = max(ans, s[n])
			} else if a > b { // 前i个物品包含当前物品
				ans = max(ans, s[i]+int64(b)*int64(n-i))
			} else { // 前i个物品不包含当前物品
				ans = max(ans, s[i]+int64(a)+int64(b)*int64(n-i-1))
			}
		}
		Fprintln(out, ans)
	}
}

func max(x, y int64) int64 {
	if x > y {
		return x
	} else {
		return y
	}
}
