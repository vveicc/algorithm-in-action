package main

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	// 每个连续段的按位或值及连续段的左端点
	type info struct{ or, i int }
	a := make([]info, n)
	for i := n - 1; i >= 0; i-- { // 枚举左端点，方便处理右端点
		x := nums[i]
		// 计算每个连续段的按位或值
		for j := range a {
			a[j].or |= x
		}
		// 增加新的连续段
		a = append(a, info{x, i})
		// 原地去重
		j := 0
		for _, y := range a[1:] {
			if y.or == a[j].or {
				a[j].i = y.i
			} else {
				j++
				a[j] = y
			}
		}
		a = a[:j+1]
		// 计算以nums[i]为左端点的按位或值最大且长度最小的子数组的长度
		ans[i] = a[0].i - i + 1
	}
	return ans
}
