package main

import "fmt"

func maxProductN(nums []int) int {
	var max int
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum *= nums[j]
		}
		fmt.Println(sum)
		if max < sum {
			max = sum
		}
	}
	return max
}
