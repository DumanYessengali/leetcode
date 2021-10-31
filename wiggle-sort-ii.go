package main

import "fmt"

func wiggleSort(nums []int) {
	sort(nums)
	//n1 := nums[:len(nums)-len(nums)/2]
	//
	//n2 := nums[len(nums)-len(nums)/2:]
	//
	//fmt.Println(n1, n2)

	k := len(nums) - len(nums)/2

	for i := 1; i < len(nums); i += 2 {
		m := k
		nums[i], nums[k] = nums[k], nums[i]
		for i+1 < m {
			nums[m], nums[m-1] = nums[m-1], nums[m]
			m--
		}
		k++
	}
	fmt.Println(nums)
}
