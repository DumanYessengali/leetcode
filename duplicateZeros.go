package main

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			arrChange := make([]int, len(arr[i+1:len(arr)-1]))
			copy(arrChange, arr[i+1:len(arr)-1])
			arr[i+1] = 0

			k := i + 2
			for j := 0; j < len(arrChange); j++ {
				arr[k] = arrChange[j]
				k++
			}
			i++
		}
	}
}
