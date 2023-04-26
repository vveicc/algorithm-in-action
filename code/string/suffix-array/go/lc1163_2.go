package main

func lastSubstring(s string) string {
	j := 0
	for i, c := range s {
		if byte(c) > s[j] {
			j = i
		}
	}
	s = s[j:]
	n := len(s)
	z := make([]int, n)
	z[0] = n
	for i, l, r := 1, 0, 0; i < n; i++ {
		z[i] = max(0, min(z[i-l], r-i+1))
		for ; i+z[i] < n && s[z[i]] == s[i+z[i]]; z[i]++ {
		}
		if j = i + z[i] - 1; j > r {
			l, r = i, j
		}
		if z[i] > 0 && i+z[i] < n && s[i+z[i]] > s[z[i]] {
			s = s[i:]
			n = len(s)
			i, l, r = 1, 0, 0
		}
	}
	return s
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
