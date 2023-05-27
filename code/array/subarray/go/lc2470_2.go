package main

func subarrayLCM(nums []int, k int) (ans int) {
	// 每个连续段的LCM及连续段的右端点
	type info struct{ lcm, i int }
	var a []info
	i0 := -1 // 最左一个连续段的左端点（不包含）
	for i, x := range nums {
		// 保证计算的LCM都是K的因子
		if k%x != 0 {
			i0, a = i, nil
			continue
		}
		// 计算每个连续段的LCM
		for j, y := range a {
			a[j].lcm *= x / gcd(y.lcm, x)
		}
		// 增加新的连续段
		a = append(a, info{x, i})
		// 原地去重
		j := 0
		for _, y := range a[1:] {
			if y.lcm == a[j].lcm {
				a[j].i = y.i
			} else {
				j++
				a[j] = y
			}
		}
		a = a[:j+1]
		// 统计以nums[i]为右端点的LCM为K的子数组的数目
		if a[0].lcm == k {
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
