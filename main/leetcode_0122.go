package main

func maxProfit2(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	maxProfit := 0
	var deltaProfit int

	for day := 1; day < len(prices); day++ {
		deltaProfit = prices[day] - prices[day-1]

		if deltaProfit < 0 {
			continue
		}
		maxProfit += deltaProfit
	}

	return maxProfit
}
