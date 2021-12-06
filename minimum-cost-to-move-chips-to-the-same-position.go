package main

func minCostToMoveChips(position []int) int {
	var even, odd int
	for _, v1 := range position {
		if v1%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return min(even, odd)
}
