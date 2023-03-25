package main

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	r := n - 1
	for ; r != 0 && arr[r-1] <= arr[r]; r-- {
	}
	if r == 0 {
		return 0 // arr已经非递减
	} else {
		ans := r            // 删除区间 [0, r)
		for l := 0; ; r++ { // 枚举右端点
			for ; r == n || arr[l] <= arr[r]; l++ { // 移动左端点
				ans = min(ans, r-l-1) // 删除区间 (l, r)
				if arr[l] > arr[l+1] {
					return ans
				}
			}
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
