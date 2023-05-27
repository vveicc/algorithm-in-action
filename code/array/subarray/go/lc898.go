package main

func subarrayBitwiseORs(arr []int) int {
	set := make(map[int]struct{})
	var ors []int // 每个连续段的按位或值
	for _, x := range arr {
		// 计算每个连续段的按位或值
		for j := range ors {
			ors[j] |= x
		}
		// 增加新的连续段
		ors = append(ors, x)
		// 原地去重
		j := 0
		set[ors[0]] = struct{}{}
		for _, or := range ors[1:] {
			if or != ors[j] {
				j++
				ors[j] = or
				set[or] = struct{}{}
			}
		}
		ors = ors[:j+1]
	}
	return len(set)
}
