package main

func countSubstrings(s string, t string) (ans int) {
	m, n := len(s), len(t)
	// i表示s子串的结束索引，j表示t子串的结束索引，k表示s子串的起始索引
	// 枚举 d = i - j ，则 j = i - d
	for d := 1 - n; d < m; d++ {
		i := max(d, 0)
		// 枚举s子串的结束索引
		// k0表示s子串结束索引i往左第二个不同字符的索引
		// k1表示s子串结束索引i往左第一个不同字符的索引
		for k0, k1 := i-1, i-1; i < m && i-d < n; i++ {
			if s[i] != t[i-d] {
				k0, k1 = k1, i
			}
			ans += k1 - k0 // (k0, k1]区间内的k都可以作为s子串的起始索引
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
