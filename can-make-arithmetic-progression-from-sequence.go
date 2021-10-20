package main

func canMakeArithmeticProgression(arr []int) bool {
	arr = sort(arr)
	d := arr[1] - arr[0]
	if arr[0] == arr[1] {
		return false
	}
	for i := 1; i < len(arr); i++ {
		if arr[i-1]+d != arr[i] {
			return false
		}
	}
	return true
}
