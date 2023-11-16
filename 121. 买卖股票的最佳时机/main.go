package main

func maxProfit(prices []int) int {
	min, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i] > min {
			profit = max(prices[i]-min, profit)
		}
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
