package main

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	if stampHeight == 1 && stampWidth == 1 {
		return true
	}
	m, n := len(grid), len(grid[0])
	s := make([][]int, m+1) // grid 的二维前缀和
	s[0] = make([]int, n+1)
	d := make([][]int, m+1) // 邮票覆盖次数的二维差分
	d[0] = make([]int, n+1)
	for i, row := range grid {
		s[i+1] = make([]int, n+1)
		d[i+1] = make([]int, n+1)
		for j, v := range row {
			s[i+1][j+1] = v + s[i][j+1] + s[i+1][j] - s[i][j]
			if x, y := i+1-stampHeight, j+1-stampWidth; x >= 0 && y >= 0 {
				if s[i+1][j+1]-s[x][j+1]-s[i+1][y]+s[x][y] == 0 { // 左上(x,y)到右下(i,j)的矩形内可以贴一张邮票
					d[x][y]++
					d[x][j+1]--
					d[i+1][y]--
					d[i+1][j+1]++
				}
			}
		}
	}
	// 复用 s 计算每个格子的邮票覆盖次数
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			s[i+1][j+1] = d[i][j] + s[i+1][j] + s[i][j+1] - s[i][j]
			if s[i+1][j+1] == 0 && grid[i][j] == 0 { // (i,j)是空格子且邮票覆盖次数为0
				return false
			}
		}
	}
	return true
}
