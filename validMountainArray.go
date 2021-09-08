package main

func validMountainArray(arr []int) bool {
	var arrLeft []int
	var arrRight []int

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}
		if arr[i] < arr[i+1] {
			continue
		} else {
			arrLeft = append(arrLeft, arr[:i]...)
			arrRight = append(arrRight, arr[i:]...)
			break
		}
	}

	if len(arrLeft) == 0 || len(arrRight) == 0 {
		return false
	}

	for i := 0; i < len(arrRight)-1; i++ {
		if arrRight[i] == arrRight[i+1] || arrRight[i] < arrRight[i+1] {

			return false
		}
		if arrRight[i] > arrRight[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}
