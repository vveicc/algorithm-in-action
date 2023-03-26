package main

func maxProfitII(prices []int) int {
	n := len(prices)
	// buy: 每天结束时, 手中持有股票的累计最大收益
	// sell: 每天结束时, 手中不持有股票的累计最大收益
	buy, sell := -prices[0], 0
	for i := 1; i < n; i++ {
		buy = max(buy, sell-prices[i])
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
