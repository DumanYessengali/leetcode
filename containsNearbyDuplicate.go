package main

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && math.Abs(float64(i-j)) <= float64(k) && i != j {
				return true
			}
		}
	}
	return false
}
