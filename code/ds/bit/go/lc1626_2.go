package main

import "sort"

const shift, mask = 10, 1<<10 - 1

func bestTeamScore(scores []int, ages []int) int {
	u := 0
	for i, x := range ages {
		ages[i] = (scores[i] << shift) | x
		u = max(u, x)
	}
	sort.Ints(ages)

	t := make(BIT, u+1)
	for _, sa := range ages {
		s, a := sa>>shift, sa&mask
		t.update(a, t.query(a)+s) // 更新年龄不超过a的无矛盾球队的最高得分
	}
	return t.query(u)
}

// 树状数组
type BIT []int

// 单点更新
func (t BIT) update(i, v int) {
	for n := len(t); i < n; i += i & -i {
		t[i] = max(t[i], v)
	}
}

// 前缀查询 [1, i]
func (t BIT) query(i int) (ans int) {
	for ; i > 0; i &= i - 1 {
		ans = max(ans, t[i])
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
