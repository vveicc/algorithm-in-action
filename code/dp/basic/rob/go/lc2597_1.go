package main

func beautifulSubsets(nums []int, k int) (ans int) {
	n := len(nums)
	for i := range nums {
		nums[i] += k
	}
	// [-k, 1000+k] => [0, 1000+k<<1]
	cnt := make([]int, 1001+k<<1)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans++
		} else {
			// 不选nums[i]
			dfs(i + 1)
			// 选nums[i]
			if x := nums[i]; cnt[x-k] == 0 && cnt[x+k] == 0 {
				cnt[x]++
				dfs(i + 1)
				cnt[x]--
			}
		}
	}
	dfs(0)
	return ans - 1 // 去除空集
}
