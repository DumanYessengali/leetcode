package main

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] == arr[j]*2 && i != j {

				return true
			}
		}
	}
	return false
}
