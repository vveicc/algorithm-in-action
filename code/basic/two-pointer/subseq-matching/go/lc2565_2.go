package main

func minimumScore(s string, t string) int {
	m, n := len(s), len(t)
	suf := make([]int, m+1)
	suf[m] = n // 哨兵，防止越界
	for i, j := m-1, n-1; i >= 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		suf[i] = j + 1 // t[j+1:]是与s[i:]匹配（子序列匹配）的最长后缀
	}
	ans := suf[0]
	if ans == 0 {
		return 0 // t 是 s 的子序列
	}
	for i, j := 0, 0; i < m; i++ {
		if s[i] == t[j] { // 注意 j 一定小于 n ，因为 j = n 时，t 是 s 的子序列，前面已经判断过
			j++ // t[:j]是与s[:i+1]匹配（子序列匹配）的最长前缀
		}
		ans = min(ans, suf[i+1]-j) // 删除 t[j:suf[i+1]]
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
