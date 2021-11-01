package main

func replaceElements(arr []int) []int {
	k := -1

	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], k = k, max(k, arr[i])
	}
	return arr
}
