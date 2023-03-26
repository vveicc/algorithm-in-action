package main

func maxProfitIV(k int, prices []int) int {
	n := len(prices)
	// buy[i]表示第i次买入后的最大利润
	buy := make([]int, k+1)
	// sell[i]表示第i次卖出后的最大利润
	sell := make([]int, k+1)
	for i := 1; i <= k; i++ {
		buy[i] = -prices[0]
	}
	for day := 1; day < n; day++ {
		for i := 0; i < k; i++ {
			buy[i+1] = max(buy[i+1], sell[i]-prices[day])
			sell[i+1] = max(sell[i+1], buy[i+1]+prices[day])
		}
	}
	return sell[k]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
