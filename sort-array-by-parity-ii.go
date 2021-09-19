package main

func sortArrayByParityII(nums []int) []int {
	n1 := make([]int, len(nums)/2)
	n2 := make([]int, len(nums)/2)
	x, y := 0, 0

	for _, v1 := range nums {
		if v1%2 == 0 {
			n1[x] = v1
			x++
		} else {
			n2[y] = v1
			y++
		}
	}
	x, y = 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = n1[x]
			x++
		} else {
			nums[i] = n2[y]
			y++
		}
	}
	return nums
}
