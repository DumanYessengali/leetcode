package main

func diagonalSum(mat [][]int) int {
	sum := 0

	n := len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == n-i-1 || i == j {
				sum += mat[i][j]
			}
		}
	}

	return sum
}
