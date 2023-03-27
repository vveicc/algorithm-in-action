package main

func countSubstrings(s string, t string) (ans int) {
	m, n := len(s), len(t)
	// l[i][j]表示s[i]和t[j]左侧（不包含）连续相同字符的个数
	l := make([][]int, m+1)
	l[0] = make([]int, n+1)
	for i, a := range s {
		l[i+1] = make([]int, n+1)
		for j, b := range t {
			if a == b {
				l[i+1][j+1] = l[i][j] + 1
			}
		}
	}
	// r[i][j]表示s[i]和t[j]右侧（包含）连续相同字符的个数
	r := make([][]int, m+1)
	r[m] = make([]int, n+1)
	for i := m - 1; i >= 0; i-- {
		r[i] = make([]int, n+1)
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				r[i][j] = r[i+1][j+1] + 1
			}
		}
	}
	// 枚举不同字符
	for i, a := range s {
		for j, b := range t {
			if a != b {
				// 累加以a和b为不同字符的子串对的个数
				ans += (l[i][j] + 1) * (r[i+1][j+1] + 1)
			}
		}
	}
	return
}
