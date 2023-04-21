package main

func minImpossibleOR(nums []int) int {
	or := 0
	for _, x := range nums {
		if x&(x-1) == 0 {
			or |= x
		}
	}
	or = ^or // ~x = -(x+1)
	return or & -or
}
