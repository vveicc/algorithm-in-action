package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	Fscan(in, &n, &m)

	// 所有课程的依赖关系组成森林结构
	// 新增一门编号为0且学分为0的课程作为根节点，将依赖关系转为一棵树
	// 从编号从0开始的 n+1 门课程中选择 m+1 门课程
	n++
	m++

	s := make([]int, n)
	og := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &k, &s[i])
		og[k] = append(og[k], i)
	}

	// dp[x][i]表示选择以x为根的子树中的i门课程所能获得的最大学分
	dp := make([][]int, n)
	var dfs func(x int) int
	dfs = func(x int) int {
		cx := 1
		dp[x] = make([]int, m+1)
		dp[x][1] = s[x] // 仅选择课程x
		for _, y := range og[x] {
			cy := dfs(y)
			for i := min(m, cx); i != 0; i-- { // 从已合并过的子树中选择i门课程
				for j := 1; j <= cy && i+j <= m; j++ { // 从子树y中选择j门课程
					dp[x][i+j] = max(dp[x][i+j], dp[x][i]+dp[y][j])
				}
			}
			cx += cy
		}
		return cx
	}
	dfs(0)

	Println(dp[0][min(n, m)])
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
