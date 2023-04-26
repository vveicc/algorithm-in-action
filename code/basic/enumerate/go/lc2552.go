package main

func countQuadruplets(nums []int) int64 {
	ans := 0
	n := len(nums)
	// cnt[j]表示前序满足条件的三元组(i, j, k)的个数
	cnt := make([]int, n)
	for q, x := range nums {
		ltk := 0
		for p, y := range nums[:q] { // p < q
			if x > y {
				ans += cnt[p] // 计算合法四元组(_, j=p, _, l=q)的个数
				ltk++         // 计算满足i=p<q=k且nums[i]=y<x=nums[k]的二元组(i=p, k=q)的个数
			} else {
				cnt[p] += ltk // 计算满足i<j=p<q=k且nums[i]<nums[k]=x<y=nums[j]的三元组(_, j=p, k=q)的个数
			}
		}
	}
	return int64(ans)
}
