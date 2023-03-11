package main

func countTriplets(nums []int) (ans int) {
	n := len(nums)
	cnt := [1 << 16]int{n}
	for _, x := range nums {
		x ^= 0xffff
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
