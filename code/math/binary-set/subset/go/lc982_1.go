package main

func countTriplets(nums []int) (ans int) {
	cnt := [1 << 16]int{}
	for _, x := range nums {
		for _, y := range nums {
			cnt[x&y]++
		}
	}
	for a, c := range cnt {
		for _, z := range nums {
			if a&z == 0 {
				ans += c
			}
		}
	}
	return
}
