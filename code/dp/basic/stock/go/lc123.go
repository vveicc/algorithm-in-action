package main

func maxProfitIII(prices []int) int {
	n := len(prices)
	// 第一次买、卖后的最大利润
	buy1, sell1 := -prices[0], 0
	// 第二次买、卖后的最大利润
	buy2, sell2 := -prices[0], 0
	for i := 1; i < n; i++ {
		// 使第一次买、卖后的利润最大化
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		// 使第二次买、卖后的利润最大化
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
