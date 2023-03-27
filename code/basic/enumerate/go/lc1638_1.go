package main

func countSubstrings(s string, t string) (ans int) {
	m, n := len(s), len(t)
	for i := 0; i < m; i++ { // 枚举s子串的起始索引
		for j := 0; j < n; j++ { // 枚举t子串的起始索引
			for k, d := 0, 0; i+k < m && j+k < n && d < 2; k++ { // 枚举子串长度
				if s[i+k] != t[j+k] {
					d++
				}
				if d == 1 {
					ans++
				}
			}
		}
	}
	return
}
