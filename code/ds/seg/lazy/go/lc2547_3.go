package main

func minCost(nums []int, k int) (ans int) {
	n := len(nums)
	pre := make([]int, n)
	pre2 := make([]int, n)
	t := buildST(n)
	for i, x := range nums {
		i++                 // 线段树区间从1开始
		t.update(i, i, ans) // 设置前一个f[i+1]的值
		t.update(pre[x]+1, i, -1)
		if pre[x] != 0 { // x不是首次出现
			t.update(pre2[x]+1, pre[x], 1)
		}
		ans = k + t.query(1, i)
		pre2[x] = pre[x]
		pre[x] = i
	}
	return ans + n
}

func buildST(n int) (t seg) {
	t = make(seg, n<<2)
	var dfs func(o, l, r int)
	dfs = func(o, l, r int) {
		t[o].l, t[o].r = l, r
		if l != r {
			mid, lo := t.mid(o), o<<1
			dfs(lo, l, mid)
			dfs(lo|1, mid+1, r)
		}
	}
	dfs(1, 1, n)
	return
}

type seg []struct{ l, r, min, add int }

func (t seg) mid(o int) int  { return (t[o].l + t[o].r) >> 1 }
func (t seg) maintain(o int) { t[o].min = min(t[o<<1].min, t[o<<1|1].min) }

func (t seg) do(o, x int) {
	t[o].min += x
	t[o].add += x
}

func (t seg) spread(o int) {
	if t[o].add != 0 {
		t.do(o<<1, t[o].add)
		t.do(o<<1|1, t[o].add)
		t[o].add = 0
	}
}

func (t seg) update(l, r, x int) {
	var dfs func(o int)
	dfs = func(o int) {
		if l <= t[o].l && t[o].r <= r {
			t.do(o, x)
		} else {
			t.spread(o)
			mid, lo := t.mid(o), o<<1
			if l <= mid {
				dfs(lo)
			}
			if mid < r {
				dfs(lo | 1)
			}
			t.maintain(o)
		}
	}
	dfs(1)
}

func (t seg) query(l, r int) int {
	var dfs func(o int) int
	dfs = func(o int) int {
		if l <= t[o].l && t[o].r <= r {
			return t[o].min
		} else {
			t.spread(o)
			mid, lo := t.mid(o), o<<1
			if r <= mid {
				return dfs(lo)
			} else if l > mid {
				return dfs(lo | 1)
			} else {
				return min(dfs(lo), dfs(lo|1))
			}
		}
	}
	return dfs(1)
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
