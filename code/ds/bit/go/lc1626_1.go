package main

import "sort"

const shift, mask = 32, 1<<32 - 1

func bestTeamScore(scores []int, ages []int) (ans int) {
	n := len(scores)
	for i, x := range scores {
		scores[i] = (ages[i] << shift) | x
	}
	sort.Ints(scores)

	// f[i]表示以scores[i]结尾的无矛盾球队的最高得分
	f := make([]int, n)
	for i := range scores {
		scores[i] &= mask
		for j, score := range scores[:i] {
			if score <= scores[i] {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] += scores[i]
		ans = max(ans, f[i])
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
