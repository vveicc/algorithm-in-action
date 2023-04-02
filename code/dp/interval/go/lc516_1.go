package main

func longestPalindromeSubseq(s string) int {
	n := len(s)
	f := make([]int, n+1)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i] == s[n-1-j] { // t[j] = s[n-1-j]
				f[j+1] = pre[j] + 1
			} else {
				f[j+1] = max(pre[j+1], f[j])
			}
		}
		f, pre = pre, f // 滚动数组优化
	}
	return pre[n]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
