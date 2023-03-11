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

	var n, m, d, u, v int
	Fscan(in, &n, &m, &d)
	monster := make([]bool, n+1)
	for i := 0; i < m; i++ {
		Fscan(in, &u)
		monster[u] = true
	}
	ug := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &u, &v)
		ug[u] = append(ug[u], v)
		ug[v] = append(ug[v], u)
	}

	const inf = int(-1e9)

	// 第一遍DFS以节点1为根，计算每个节点与其子树中的最远怪物距离、次远怪物距离及最远怪物距离所在子树节点
	type tuple struct{ f, s, o int }
	dis := make([]tuple, n+1)
	var dfs1 func(x, fa int) int
	dfs1 = func(x, fa int) int {
		// 距离初始化为inf，方便处理子树中没有怪物的情况
		f, s, o := inf, inf, 0
		for _, y := range ug[x] {
			if y != fa {
				if yf := dfs1(y, x) + 1; yf > f {
					f, s, o = yf, f, y
				} else if yf > s {
					s = yf
				}
			}
		}
		dis[x] = tuple{f, s, o}
		if f < 0 && monster[x] {
			return 0
		} else {
			return f
		}
	}
	dfs1(1, 0)

	// 第二遍DFS通过换根DP计算每个节点为根时，与其子树中的最远怪物距离，判断是否可以作为传送门
	cnt := 0
	var dfs func(x, fa, df int)
	dfs = func(x, fa, df int) {
		if df <= d {
			dx := dis[x]
			if dx.f <= d {
				cnt++
			}
			if df < 0 && monster[x] {
				df = 0
			}
			for _, y := range ug[x] {
				if y != fa {
					if y == dx.o {
						dfs(y, x, max(df, dx.s)+1)
					} else {
						dfs(y, x, max(df, dx.f)+1)
					}
				}
			}
		}
	}
	dfs(1, 0, inf)
	Fprintln(out, cnt)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
