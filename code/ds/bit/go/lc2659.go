package main

import "sort"

func countOperationsToEmptyArray(nums []int) int64 {
	ans := 0
	n := len(nums)
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return nums[id[i]] < nums[id[j]] })

	// [-1, n-1] -> [0, n]
	t := make(BIT, n+1)
	pre := 0
	for _, i := range id {
		i += 1
		if i > pre {
			ans += i - pre - (t.query(i) - t.query(pre))
		} else {
			ans += n - (pre - i) - (t.query(n) - t.query(pre) + t.query(i))
		}
		t.update(i, 1)
		pre = i
	}
	return int64(ans)
}

type BIT []int

// 单点更新
func (t BIT) update(i, val int) {
	for n := len(t); i < n; i += i & -i {
		t[i] += val
	}
}

// 查询前缀区间[1, i]
func (t BIT) query(i int) (ans int) {
	for ; i > 0; i &= i - 1 {
		ans += t[i]
	}
	return
}
