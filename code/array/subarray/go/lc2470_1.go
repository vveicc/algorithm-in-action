package main

func subarrayLCM(nums []int, k int) (ans int) {
	for i := range nums {
		lcm := 1
		for _, x := range nums[i:] {
			if lcm *= x / gcd(lcm, x); lcm == k {
				ans++
			} else if k%lcm != 0 {
				break
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
