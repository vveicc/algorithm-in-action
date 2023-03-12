package main

import (
	"sort"
	"strings"
)

func braceExpansionII(expression string) []string {
	set := make(map[string]struct{})
	var dfs func(expr string)
	dfs = func(expr string) {
		r := strings.IndexByte(expr, '}')
		if r == -1 {
			set[expr] = struct{}{}
		} else {
			l := r - 1
			for ; expr[l] != '{'; l-- {
			}
			for _, s := range strings.Split(expr[l+1:r], ",") {
				dfs(expr[:l] + s + expr[r+1:])
			}
		}
	}
	dfs(expression)

	ans := make([]string, 0, len(set))
	for s := range set {
		ans = append(ans, s)
	}
	sort.Strings(ans)
	return ans
}
