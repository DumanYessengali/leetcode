package main

func minTimeToVisitAllPoints(points [][]int) int {
	k := 0
	for i, j := 0, 1; j < len(points); i, j = i+1, j+1 {
		k += max(abs(points[j][0]-points[i][0]), abs(points[j][1]-points[i][1]))
	}
	return k
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
