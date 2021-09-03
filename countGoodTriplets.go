package main

import "math"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	num := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				a1 := arr[i]
				a2 := arr[j]
				a3 := arr[k]
				if isTrue(a1, a2, a) && isTrue(a2, a3, b) && isTrue(a1, a3, c) {
					num++
				}
			}
		}
	}

	return num
}

func isTrue(n1, n2, num int) bool {
	if int(math.Abs(float64(n1-n2))) <= num {
		return true
	}

	return false
}
