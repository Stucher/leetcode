package main

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	maxProfit := 0
	dp := make([]int, len(prices))
	var tmpProfit int

	for day := 0; day < len(prices); day++ {
		if day == 0 {
			dp[day] = 0
			continue
		}

		if tmpProfit = prices[day] - prices[day-1] + dp[day-1]; tmpProfit > 0 {
			dp[day] = tmpProfit
		} else {
			dp[day] = 0
		}

		if dp[day] > maxProfit {
			maxProfit = dp[day]
		}
	}

	return maxProfit
}
