package main

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	for c := 0; c < n; c++ {
		sortByDiagonal(mat, 0, c, n, m)
	}
	for r := 0; r < m; r++ {
		sortByDiagonal(mat, r, 0, n, m)
	}

	return mat
}

func sortByDiagonal(mat [][]int, row, col, n, m int) {
	r, c := row, col
	var arr []int
	for r < m && c < n {
		arr = append(arr, mat[r][c])
		r++
		c++
	}

	arr = mergeS(arr)
	r, c = row, col
	for i := range arr {
		mat[r][c] = arr[i]
		r++
		c++
	}
}
