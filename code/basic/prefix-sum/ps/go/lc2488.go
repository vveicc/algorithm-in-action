package main

func countSubarrays(nums []int, k int) (ans int) {
	j := 0
	for ; nums[j] != k; j++ {
	}
	s := len(nums)
	// [-n, n] => [0, n<<1]
	cnt := make([]int, s<<1+1)
	cnt[s] = 1
	for i, x := range nums {
		if x > k {
			s++
		} else {
			s--
		}
		if i < j {
			cnt[s]++
		} else {
			ans += cnt[s] + cnt[s+1]
		}
	}
	return
}
