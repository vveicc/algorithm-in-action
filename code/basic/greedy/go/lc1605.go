package main

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i, j := 0, 0; i < m && j < n; {
		x := min(rowSum[i], colSum[j])
		ans[i][j] = x // 取不超过rowSum[i]和colSum[j]的最大值
		if rowSum[i] -= x; rowSum[i] == 0 {
			i++ // 该行剩余元素赋0
		}
		if colSum[j] -= x; colSum[j] == 0 {
			j++ // 该列剩余元素赋0
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
