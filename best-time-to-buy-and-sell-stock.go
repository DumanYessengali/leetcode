package main

import "fmt"

func maxProfit(prices []int) int {
	max := -1
	min := 10000
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if prices[i]-min > max {
			fmt.Println()
			max = prices[i] - min
		}
	}
	return max
}
