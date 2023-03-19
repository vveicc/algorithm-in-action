package main

func checkPalindromeFormation(a string, b string) bool {
	n := len(a)
	check := func(s, t string) bool {
		l, r := (n-1)>>1, n>>1
		for ; r < n && s[l] == s[r]; r++ {
			l--
		}
		if r == n {
			return true // s回文
		} else {
			checkLR := func(pref, suff string) bool {
				i, j := l, r
				for ; j < n && pref[i] == suff[j]; j++ {
					i--
				}
				return j == n
			}
			// s[:r]+t[r:]构成回文 || t[:l+1]+s[l+1:]构成回文
			return checkLR(s, t) || checkLR(t, s)
		}
	}
	return check(a, b) || check(b, a)
}
