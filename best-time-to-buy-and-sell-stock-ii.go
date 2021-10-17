package main

func maxProfitII(prices []int) int {
	var max int
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			max = max + prices[i] - prices[i-1]
		}
	}
	return max
}
