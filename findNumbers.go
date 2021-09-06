package main

import "strconv"

func findNumbers(nums []int) int {
	count := 0

	for _, v1 := range nums {
		k := strconv.Itoa(v1)
		if len(k)%2 == 0 {
			count++
		}
	}

	return count
}
