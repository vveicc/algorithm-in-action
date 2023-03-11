package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, l, r int
	Fscan(in, &n)
	type tuple struct{ l, r, i int }
	arr := make([]tuple, n)
	rs := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &l, &r)
		arr[i] = tuple{l, r, i}
		rs[i] = r
	}

	// 按照左端点降序排序
	sort.Slice(arr, func(i, j int) bool { return arr[i].l > arr[j].l })

	ans := make([]int, n)
	// 离散化树状数组计算前序有多少个区间右端点小于当前区间右端点
	sort.Ints(rs)
	t := make(BIT, n+1)
	for _, lri := range arr {
		x := sort.SearchInts(rs, lri.r)
		ans[lri.i] = t.query(x)
		t.update(x+1, 1)
	}

	for _, x := range ans {
		Fprintln(out, x)
	}
}

// 树状数组
type BIT []int

// 单点更新
func (t BIT) update(i, v int) {
	for n := len(t); i < n; i += i & -i {
		t[i] += v
	}
}

// 前缀查询 [1, i]
func (t BIT) query(i int) (ans int) {
	for ; i > 0; i &= i - 1 {
		ans += t[i]
	}
	return
}
