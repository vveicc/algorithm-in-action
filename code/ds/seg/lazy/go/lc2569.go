package main

func handleQuery(nums1 []int, nums2 []int, queries [][]int) (ans []int64) {
	s := 0
	for _, x := range nums2 {
		s += x
	}
	t := buildST(nums1)
	for _, q := range queries {
		if q[0] == 1 {
			t.update(q[1]+1, q[2]+1)
		} else if q[0] == 2 {
			s += q[1] * t[1].c1
		} else {
			ans = append(ans, int64(s))
		}
	}
	return
}

func buildST(arr []int) (t seg) {
	n := len(arr)
	t = make(seg, n<<2)
	var dfs func(o, l, r int)
	dfs = func(o, l, r int) {
		t[o].l, t[o].r = l, r
		if l == r {
			t[o].c1 = arr[l-1]
		} else {
			mid, lo := t.mid(o), o<<1
			dfs(lo, l, mid)
			dfs(lo|1, mid+1, r)
			t.maintain(o)
		}
	}
	dfs(1, 1, n)
	return
}

type seg []struct {
	l, r, c1 int
	flip     bool // 01翻转懒惰标记
}

func (t seg) mid(o int) int  { return (t[o].l + t[o].r) >> 1 }
func (t seg) maintain(o int) { t[o].c1 = t[o<<1].c1 + t[o<<1|1].c1 }

func (t seg) do(o int) {
	t[o].c1 = t[o].r - t[o].l + 1 - t[o].c1 // 更新节点o区间维护的信息
	t[o].flip = !t[o].flip                  // 懒惰标记，暂不下钻更新子区间，需要用到子区间时再下钻更新子区间
}

func (t seg) spread(o int) {
	if t[o].flip { // 下钻更新子区间
		t.do(o << 1)
		t.do(o<<1 | 1)
		t[o].flip = false
	}
}

func (t seg) update(l, r int) {
	var dfs func(o int)
	dfs = func(o int) {
		if l <= t[o].l && t[o].r <= r {
			t.do(o)
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
