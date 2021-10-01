package main

import "fmt"

func containsPattern(arr []int, m int, k int) bool {
	for i := 0; i < len(arr)-m; i++ {
		if len(arr) < 3 {
			if checkIsArrayEqual(arr[i:m+i], arr[i+1:]) {
				k--
			}
			if k < 2 {
				return true
			}
		}
		for j := i; j < len(arr)-m; j++ {
			var arr2 []int
			p := 0

			for y := 0; y < k; y++ {
				if p == len(arr[i:m+i]) {
					p = 0
				}
				arr2 = append(arr2, arr[i:m+i]...)
				p++
			}
			fmt.Println(arr[i:len(arr[i:m+i])*k], arr2)
			if checkIsArrayEqual(arr[i:len(arr[i:m+i])*k], arr2) {
				return true
			}
		}
	}

	return false
}

func checkIsArrayEqual(num1, num2 []int) bool {
	for i := range num1 {
		if num1[i] != num2[i] {
			return false
		}
	}
	return true
}
