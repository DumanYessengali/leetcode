package main

func peakIndexInMountainArray(arr []int) int {
	for i := range arr {
		if arr[i] > arr[i+1] {
			return i
		}
	}
	return 0
}
