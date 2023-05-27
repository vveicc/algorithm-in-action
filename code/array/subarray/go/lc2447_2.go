package main

func subarrayGCD(nums []int, k int) (ans int) {
	// 每个连续段的GCD及连续段的右端点
	type info struct{ g, i int }
	var a []info
	i0 := -1 // 最左一个连续段的左端点（不包含）
	for i, x := range nums {
		// 保证计算的GCD都是K的倍数
		if x%k != 0 {
			i0, a = i, nil
			continue
		}
		// 计算每个连续段的GCD
		for j, y := range a {
			a[j].g = gcd(x, y.g)
		}
		// 增加新的连续段
		a = append(a, info{x, i})
		// 原地去重
		j := 0
		for _, y := range a[1:] {
			if y.g == a[j].g {
				a[j].i = y.i
			} else {
				j++
				a[j] = y
			}
		}
		a = a[:j+1]
		// 统计以nums[i]为右端点的GCD为K的子数组的数目
		if a[0].g == k {
			ans += a[0].i - i0
		}
	}
	return
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
