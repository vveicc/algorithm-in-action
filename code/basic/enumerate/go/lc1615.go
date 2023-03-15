package main

func maximalNetworkRank(n int, roads [][]int) (ans int) {
	degrees := make([]int, n)
	connect := make([][]int, n)
	for i := range connect {
		connect[i] = make([]int, n)
	}
	for _, road := range roads {
		x, y := road[0], road[1]
		connect[x][y] = 1
		connect[y][x] = 1
		degrees[x]++
		degrees[y]++
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank := degrees[i] + degrees[j] - connect[i][j]
			ans = max(ans, rank)
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
