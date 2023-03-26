package main

func maxProfitII(prices []int) (ans int) {
	n := len(prices)
	for i := 1; i < n; i++ {
		// 贪心，能赚就赚
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
