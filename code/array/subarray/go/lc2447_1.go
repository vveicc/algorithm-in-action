package main

func subarrayGCD(nums []int, k int) (ans int) {
	for i, x := range nums {
		for _, y := range nums[i:] {
			if x = gcd(x, y); x == k {
				ans++
			}
		}
	}
	return
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
