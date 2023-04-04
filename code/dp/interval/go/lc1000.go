package main

func mergeStones(stones []int, k int) int {
	if n := len(stones); (n-1)%(k-1) != 0 {
		return -1
	} else {
		s := make([]int, n+1)
		for i, x := range stones {
			s[i+1] = s[i] + x
		}
		f := make([][]int, n)
		for i := n - 1; i >= 0; i-- {
			f[i] = make([]int, n)
			for j := i + 1; j < n; j++ {
				f[i][j] = 1e8
				for x := i; x < j; x += k - 1 {
					f[i][j] = min(f[i][j], f[i][x]+f[x+1][j])
				}
				if (j-i)%(k-1) == 0 {
					f[i][j] += s[j+1] - s[i] // 合并成一堆
				}
			}
		}
		return f[0][n-1]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
