package main

func lastSubstring(s string) string {
	n := len(s)
	i, j, k := 0, 1, 0
	for {
		for k = 0; j+k < n && s[i+k] == s[j+k]; k++ {
		}
		if j+k == n {
			return s[i:]
		} else if s[i+k] < s[j+k] {
			i, j = j, max(j+1, i+k+1)
		} else {
			j = j + k + 1
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
