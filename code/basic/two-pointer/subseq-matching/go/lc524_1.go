package main

func findLongestWord(s string, dictionary []string) (ans string) {
	m := len(s)
	for _, t := range dictionary {
		i, j, n := 0, 0, len(t)
		for ; i < m && j < n; i++ {
			if s[i] == t[j] {
				j++
			}
		}
		if j == n && (n > len(ans) || (n == len(ans) && t < ans)) {
			ans = t
		}
	}
	return
}
