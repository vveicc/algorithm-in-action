package main

import "sort"

func putMarbles(weights []int, k int) int64 {
	if n := len(weights); n == k {
		return 0
	} else {
		for i, x := range weights[1:] {
			weights[i] += x
		}
		n--
		sort.Ints(weights[:n])
		if k--; k > n>>1 {
			k = n - k
		}
		ans := 0
		for i, x := range weights[n-k : n] {
			ans += x - weights[i]
		}
		return int64(ans)
	}
}
