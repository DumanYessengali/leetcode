package main

import "fmt"

func threeSum(nums []int) [][]int {
	var arr [][]int

	if len(nums) < 3 {
		return [][]int{}
	}
	nums = sort(nums)
	k := 0
	j := len(nums) / 3
	l := (len(nums) / 3) * 2
	for l <= len(nums)-1 {
		fmt.Println()
		if nums[k]+nums[j]+nums[l] == 0 {

			if len(arr) == 0 {
				arr = append(arr, []int{nums[k], nums[j], nums[l]})
			} else {
				if checkArr(arr, []int{nums[k], nums[j], nums[l]}) {
					arr = append(arr, []int{nums[k], nums[j], nums[l]})
				}
			}
		}
		if k < len(nums)/3-1 {
			k++
		} else {
			if j < (len(nums)/3)*2-1 {
				j++
				k = 0
			} else {
				j = len(nums) / 3
				l++
			}
		}

	}
	return arr
}

func sort(arr []int) []int {
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func checkArr(arr [][]int, nums []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i][0] == nums[0] && arr[i][1] == nums[1] && arr[i][2] == nums[2] {
			return false
		}
	}
	return true
}
