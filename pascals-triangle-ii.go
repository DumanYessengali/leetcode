package main

func getRow(rowIndex int) []int {
	return generatePascal(insertZeroToArray(1, make([][]int, rowIndex+1)), 0, 0)[rowIndex]
}
