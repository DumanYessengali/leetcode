package main

func generate(numRows int) [][]int {
	return generatePascal(insertZeroToArray(1, make([][]int, numRows)), 0, 0)
}

func insertZeroToArray(i int, arrPascal [][]int) [][]int {
	if i <= len(arrPascal) {
		arr := make([]int, i)
		arrPascal[i-1] = arr
		return insertZeroToArray(i+1, arrPascal)
	}
	return arrPascal
}

func generatePascal(arrPascal [][]int, i, j int) [][]int {
	if i < len(arrPascal) {
		if j < len(arrPascal[i]) {
			if i == j || j == 0 {
				arrPascal[i][j] = 1
			} else {
				arrPascal[i][j] = arrPascal[i-1][j-1] + arrPascal[i-1][j]
			}
			return generatePascal(arrPascal, i, j+1)
		}
		return generatePascal(arrPascal, i+1, 0)
	}
	return arrPascal
}
