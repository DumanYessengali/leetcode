package main

func countNegatives(grid [][]int) int {
	var k int

	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] >= 0 {
				break
			} else {
				k++
			}
		}
	}
	return k
}
