package main

func maxProfit(prices []int) int {
	n := len(prices)
	// buy: 每天结束时, 持有股票（买入后未卖出）的最大收益
	// sell: 每天结束时, 不持有股票（未买入或卖出）的最大收益
	buy, sell := -prices[0], 0
	for i := 1; i < n; i++ {
		buy = max(buy, -prices[i])
		sell = max(sell, buy+prices[i])
	}
	return sell
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
