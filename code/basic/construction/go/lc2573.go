package main

import "bytes"

func findTheString(lcp [][]int) string {
	n := len(lcp)
	// 构造
	s := make([]byte, n)
	for c := byte('a'); c <= 'z'; c++ {
		if i := bytes.IndexByte(s, 0); i == -1 {
			break // 构造完成
		} else {
			for j := i; j < n; j++ {
				if lcp[i][j] != 0 {
					s[j] = c
				}
			}
		}
	}

	if bytes.IndexByte(s, 0) != -1 {
		return "" // 无法构造完成
	}

	// 验证
	// 如果s[i] == s[j], lcp[i][j] = lcp[i+1][j+1] + 1; 如果s[i] != s[j], lcp[i][j] = 0
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			expect := 0
			if s[i] == s[j] {
				expect = 1
				if i+1 < n && j+1 < n {
					expect += lcp[i+1][j+1]
				}
			}
			if lcp[i][j] != expect {
				return "" // 出现冲突
			}
		}
	}

	return string(s)
}
