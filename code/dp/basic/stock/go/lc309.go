package main

func maxProfit(prices []int) int {
	n := len(prices)
	// buy: 每天结束时，手中持有股票的累计最大收益
	// cool: 每天结束时，手中不持有股票，且处于冷冻期中（前一天卖出了股票）的累计最大收益
	// sell: 每天结束时，手中不持有股票，且不在冷冻期中（前一天之前卖出了股票）的累计最大收益
	buy, cool, sell := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		buy, sell = max(buy, sell-prices[i]), max(sell, cool)
		cool = buy + prices[i]
	}
	return max(cool, sell)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
