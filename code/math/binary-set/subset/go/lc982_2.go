package main

func countTriplets(nums []int) (ans int) {
	n := len(nums)
	cnt := [1 << 16]int{}
	for _, x := range nums {
		for _, y := range nums {
			cnt[x&y]++
		}
	}
	for _, x := range nums {
		if x == 0 {
			ans += n * n
		} else {
			x ^= 0xffff // 取补集
			for s := x; s != 0; s = (s - 1) & x {
				ans += cnt[s]
			}
			ans += cnt[0] // 空集也需要统计
		}
	}
	return
}
