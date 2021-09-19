package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	return bStudent(startTime, endTime, queryTime, 0, 0)
}
func bStudent(startTime []int, endTime []int, queryTime, n, i int) int {
	if i < len(endTime) {
		if endTime[i] >= queryTime && startTime[i] <= queryTime {
			n++
		}
		return bStudent(startTime, endTime, queryTime, n, i+1)
	}
	return n
}
