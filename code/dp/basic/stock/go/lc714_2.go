package main

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	cost, profit := prices[0]+fee, 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < cost {
			// 降低持仓成本
			cost = prices[i] + fee
		} else if prices[i] > cost {
			// 贪心，能赚就赚
			profit += prices[i] - cost
			// 卖在最高点，不重复交手续费
			cost = prices[i]
		}
	}
	return profit
}
