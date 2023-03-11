package main

func countTriplets(nums []int) (ans int) {
	m := 1
	for _, x := range nums {
		for ; m <= x; m <<= 1 {
		}
	}
	cnt := make([]int, m)
	cnt[0] = len(nums)
	mask := m - 1
	for _, x := range nums {
		x ^= mask
		for s := x; s != 0; s = (s - 1) & x {
			cnt[s]++
		}
	}
	for _, x := range nums {
		for _, y := range nums {
			ans += cnt[x&y]
		}
	}
	return
}
