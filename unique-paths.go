package main

func uniquePaths(m int, n int) int {

	var arr = make([][]int, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr[i] = append(arr[i], 0)

			if i == 0 || j == 0 {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j] + arr[i][j-1]
			}
		}
	}

	return arr[m-1][n-1]
}
