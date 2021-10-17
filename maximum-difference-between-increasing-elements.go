package main

func maximumDifference(prices []int) int {
	max := -1
	min := 1000000000
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if prices[i]-min > max && prices[i] != min {
			max = prices[i] - min
		}
	}
	return max
}
