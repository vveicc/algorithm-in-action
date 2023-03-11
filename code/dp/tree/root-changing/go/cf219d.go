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

	var n, v, w int
	Fscan(in, &n)
	ug := make([][]int, n+1)
	for ; n > 1; n-- {
		Fscan(in, &v, &w)
		ug[v] = append(ug[v], w<<1)   // 最低位 0 表示正向
		ug[w] = append(ug[w], v<<1|1) // 最低位 1 表示反向
	}

	// 第一遍DFS计算以节点1为根时，需要反向的边的数量
	var dfs1 func(x, fa int) int
	dfs1 = func(x, fa int) (c int) {
		for _, y := range ug[x] {
			if y>>1 != fa {
				c += y&1 + dfs1(y>>1, x)
			}
		}
		return
	}
	min := dfs1(1, 0)

	// 第二遍DFS通过换根DP计算每个节点为根时，需要反向的边的数量
	var xs []int
	var dfs func(x, fa, c int)
	dfs = func(x, fa, c int) {
		if c < min {
			xs, min = []int{x}, c
		} else if c == min {
			xs = append(xs, x)
		}
		for _, y := range ug[x] {
			if y>>1 != fa {
				// 如果x->y为正向，则以节点y为根时，需要反向的边的数量+1；如果为反向则-1
				dfs(y>>1, x, c-(y&1)<<1+1)
			}
		}
	}
	dfs(1, 0, min)

	sort.Ints(xs)
	Fprintln(out, min)
	for _, x := range xs {
		Fprint(out, x, " ")
	}
}
