package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, op, x, y int
	var k int64
	Fscan(in, &n, &m, &p)
	arr := make([]int64, n)
	for i := range arr {
		Fscan(in, &arr[i])
		arr[i] %= p
	}
	t := buildST(arr)
	for ; m > 0; m-- {
		Fscan(in, &op, &x, &y)
		if op == 1 {
			Fscan(in, &k)
			t.update(x, y, 0, k%p)
		} else if op == 2 {
			Fscan(in, &k)
			t.update(x, y, k%p, 1)
		} else {
			Fprintln(out, t.query(x, y))
		}
	}
}

var p int64

func buildST(arr []int64) (t seg) {
	n := len(arr)
	t = make(seg, n<<2)
	var dfs func(o, l, r int)
	dfs = func(o, l, r int) {
		t[o].l, t[o].r, t[o].mul = l, r, 1
		if l == r {
			t[o].sum = arr[l-1]
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
	l, r     int
	sum      int64
	add, mul int64
}

func (t seg) mid(o int) int  { return (t[o].l + t[o].r) >> 1 }
func (t seg) maintain(o int) { t[o].sum = (t[o<<1].sum + t[o<<1|1].sum) % p }

func (t seg) do(o int, add, mul int64) {
	t[o].sum = (t[o].sum*mul + int64(t[o].r-t[o].l+1)*add) % p
	t[o].add = (t[o].add*mul + add) % p
	t[o].mul = (t[o].mul * mul) % p
}

func (t seg) spread(o int) {
	if t[o].add != 0 || t[o].mul != 1 {
		t.do(o<<1, t[o].add, t[o].mul)
		t.do(o<<1|1, t[o].add, t[o].mul)
		t[o].add, t[o].mul = 0, 1
	}
}

func (t seg) update(l, r int, add, mul int64) {
	var dfs func(o int)
	dfs = func(o int) {
		if l <= t[o].l && t[o].r <= r {
			t.do(o, add, mul)
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

func (t seg) query(l, r int) int64 {
	var dfs func(o int) int64
	dfs = func(o int) int64 {
		if l <= t[o].l && t[o].r <= r {
			return t[o].sum
		} else {
			t.spread(o)
			mid, lo := t.mid(o), o<<1
			if r <= mid {
				return dfs(lo)
			} else if mid < l {
				return dfs(lo | 1)
			} else {
				return dfs(lo) + dfs(lo|1)
			}
		}
	}
	return dfs(1) % p
}
