package main

func minSubarray(nums []int, p int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	if s %= p; s == 0 {
		return 0
	} else {
		n := len(nums)
		ans, pre := n, 0
		// last[pre]保存前缀和余p为pre的最大索引
		last := map[int]int{0: -1}
		for i, x := range nums {
			pre = (pre + x) % p
			t := (p - (s - pre)) % p
			if j, ok := last[t]; ok {
				ans = min(ans, i-j)
			}
			last[pre] = i
		}
		if ans == n {
			return -1
		} else {
			return ans
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
