package main

func longestPalindromeSubseq(s string) int {
	n := len(s)
	f := make([]int, n)
	pre := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		f[i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				f[j] = pre[j-1] + 2
			} else {
				f[j] = max(pre[j], f[j-1])
			}
		}
		f, pre = pre, f // 滚动数组优化
	}
	return pre[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
