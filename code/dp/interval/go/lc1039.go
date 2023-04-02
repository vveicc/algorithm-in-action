package main

import "math"

func minScoreTriangulation(values []int) int {
	n := len(values)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			f[i][j] = math.MaxInt
			for k := i + 1; k < j; k++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j]+values[i]*values[j]*values[k])
			}
		}
	}
	return f[0][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
