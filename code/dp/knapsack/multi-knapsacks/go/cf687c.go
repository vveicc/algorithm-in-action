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

	var n, k int
	Fscan(in, &n, &k)
	c := make([]int, n)
	for i := range c {
		Fscan(in, &c[i])
	}

	// dp[i][j1][j2]表示从前i个数中是否能选出两个不相交的子集，其中子集1的元素和为j1，子集2的元素和为j2
	dp := make([][]bool, k+1)
	for i := range dp {
		dp[i] = make([]bool, k+1)
	}
	dp[0][0] = true

	sort.Ints(c) // 排序优化枚举效率

	var s, ms int
	for _, x := range c {
		s += x
		ms = min(s, k) // 子集元素和从当前最大和开始枚举
		for j1 := ms; j1 >= 0; j1-- {
			for j2 := ms; j2 >= 0; j2-- {
				// 不选择x，或者选择x放入子集1，或者选择x放入子集2
				dp[j1][j2] = dp[j1][j2] || (j1 >= x && dp[j1-x][j2]) || (j2 >= x && dp[j1][j2-x])
			}
		}
	}

	var ans []int
	for x, fx := range dp {
		if fx[k-x] {
			ans = append(ans, x)
		}
	}
	Fprintln(out, len(ans))
	for _, x := range ans {
		Fprint(out, x, " ")
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
