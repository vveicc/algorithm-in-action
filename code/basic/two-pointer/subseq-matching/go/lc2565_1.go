package main

import "sort"

func minimumScore(s string, t string) int {
	m, n := len(s), len(t)
	lIdx := make([]int, n)
	for i, j := -1, 0; j < n; j++ {
		for i++; i < m && s[i] != t[j]; i++ {
		}
		lIdx[j] = i // s[:i+1]是与t[:j+1]匹配（子序列匹配）的最短前缀
	}
	if lIdx[n-1] < m {
		return 0
	}
	rIdx := make([]int, n)
	for i, j := m, n-1; j >= 0; j-- {
		for i--; i >= 0 && s[i] != t[j]; i-- {
		}
		rIdx[j] = i // s[i:]是与t[j:]匹配（子序列匹配）的最短后缀
	}
	return sort.Search(n, func(x int) bool {
		if rIdx[x] > -1 || lIdx[n-1-x] < m { // 删除 t[:x] 或 t[n-x:]
			return true
		} else {
			for r := x + 1; r < n; r++ {
				if rIdx[r] > lIdx[r-x-1] { // 删除 t[r-x:r]
					return true
				}
			}
			return false
		}
	})
}
