package main

func prevPermOpt1(arr []int) []int {
	n := len(arr)
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] { // 找到最大的满足条件的i，此时arr[i+1:]是不严格递增（非递减）的
			j := n - 1
			for ; arr[j] >= arr[i] || arr[j-1] == arr[j]; j-- {
			}
			arr[i], arr[j] = arr[j], arr[i]
			break
		}
	}
	return arr
}
