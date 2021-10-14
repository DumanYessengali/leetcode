package main

func mergeIntervals(intervals [][]int) [][]int {
	//should fix
	var arr [][]int
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] <= intervals[i][0] {
			arr = append(arr, []int{intervals[i-1][0], intervals[i][1]})
		}
	}
	return arr
}
