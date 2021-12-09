package main

func canReach(arr []int, start int) bool {
	for start >= 0 && start < len(arr) && arr[start] < len(arr) {
		j := arr[start]
		arr[start] += len(arr)
		return j == 0 || canReach(arr, start+j) || canReach(arr, start-j)
	}
	return false
}
