package main

func minCostClimbingStairs(cost []int) int {
	var x, y int
	for i := 1; i < len(cost); i++ {
		x, y = y, min(cost[i-1]+x, cost[i]+y)
	}
	return y
}
