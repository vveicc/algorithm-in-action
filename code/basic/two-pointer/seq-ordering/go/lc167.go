package main

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1} // 下标从 1 开始
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
}
