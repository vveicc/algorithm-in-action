package main

import (
	"sort"
	"unicode"
)

func braceExpansionII(expression string) []string {
	idx, n := 0, len(expression)

	// expr: term[,term]   每个表达式由多个term以逗号分隔
	// term: item[item]    每个term由多个item相接，无逗号分隔
	// item: letter|{expr} 每个item可以是单个小写英文字母，或者由花括号包裹的表达式
	var expr func() []string
	var term func() []string
	var item func() []string

	expr = func() (ret []string) {
		ret = append(ret, term()...)
		for idx < n && expression[idx] == ',' {
			idx++
			ret = append(ret, term()...)
		}
		return
	}

	term = func() (ret []string) {
		var tmp []string
		ret = []string{""}
		for idx < n && (expression[idx] == '{' || unicode.IsLetter(rune(expression[idx]))) {
			sub := item()
			for _, s := range ret {
				for _, t := range sub {
					tmp = append(tmp, s+t)
				}
			}
			ret, tmp = tmp, nil
		}
		return
	}

	item = func() (ret []string) {
		if expression[idx] == '{' {
			idx++
			ret = expr()
		} else {
			ret = []string{string(expression[idx])}
		}
		idx++
		return
	}

	set := make(map[string]struct{})
	for _, s := range expr() {
		set[s] = struct{}{}
	}
	ans := make([]string, 0, len(set))
	for s := range set {
		ans = append(ans, s)
	}
	sort.Strings(ans)
	return ans
}
