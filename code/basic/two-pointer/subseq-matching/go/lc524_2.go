package main

func findLongestWord(s string, dictionary []string) (ans string) {
	m := len(s)
	var first [26]int
	for i := range first {
		first[i] = m
	}
	// next[i][j]表示s[i]后第一个字母j+'a'的索引
	next := make([][]int, m)
	for i := m - 1; i >= 0; i-- {
		next[i] = make([]int, 26)
		copy(next[i], first[:])
		first[s[i]-'a'] = i
	}
	for _, t := range dictionary {
		j, n := 0, len(t)
		for i := first[t[j]-'a']; i < m; i = next[i][t[j]-'a'] {
			if j++; j == n {
				if n > len(ans) || (n == len(ans) && t < ans) {
					ans = t
				}
				break
			}
		}
	}
	return
}
