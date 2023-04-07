package main

import "sort"

const shift, mask = 5, 1<<5 - 1

func beautifulSubsets(nums []int, k int) int {
	// 同余分组
	groups := make(map[int]map[int]int)
	for _, x := range nums {
		if groups[x%k] == nil {
			groups[x%k] = make(map[int]int)
		}
		groups[x%k][x]++
	}
	ans := 1
	for _, g := range groups {
		n := len(g)
		xcs := make([]int, 0, n)
		for x, c := range g {
			xcs = append(xcs, x<<shift|c)
		}
		sort.Ints(xcs)
		// 打家劫舍
		f0, f1 := 1, 1<<(xcs[0]&mask)
		for i := 1; i < n; i++ {
			x, c := xcs[i]>>shift, xcs[i]&mask
			if x-xcs[i-1]>>shift == k {
				f0, f1 = f1, f1+f0*(1<<c-1)
			} else {
				f0, f1 = f1, f1<<c
			}
		}
		ans *= f1 // 乘法原理，不同分组的美丽子集可以任意组合
	}
	return ans - 1 // 去除空集
}
