package main

func longestWPI(hours []int) (ans int) {
	s := 0                     // 前缀和
	first := make(map[int]int) // 前缀和首次出现的索引
	for i, h := range hours {
		if h > 8 {
			s++
		} else {
			s--
		}
		if s > 0 {
			ans = i + 1
		} else if j, ok := first[s-1]; ok {
			ans = max(ans, i-j)
		}
		if _, ok := first[s]; !ok {
			first[s] = i
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
