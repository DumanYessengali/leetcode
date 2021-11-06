package main

func mergeIntervals(intervals [][]int) [][]int {
	var arr [][]int
	if len(intervals) <= 1 {
		return intervals
	}
	for i := 1; i < len(intervals); i += 2 {
		if intervals[i-1][1] >= intervals[i][0] && intervals[i][0] < intervals[i-1][0] {
			arr = append(arr, []int{intervals[i-1][0], intervals[i][1]})
		} else if intervals[i-1][1] >= intervals[i][0] && intervals[i][0] > intervals[i-1][0] {
			arr = append(arr, []int{min(intervals[i][0], intervals[i-1][0]), intervals[i][1]})
		} else {
			arr = append(arr, intervals[i-1])
			arr = append(arr, intervals[i])
		}
	}
	return arr
}
