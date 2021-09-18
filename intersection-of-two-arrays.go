package main

func intersection(nums1 []int, nums2 []int) []int {
	arrMap := map[int]bool{}
	var newNums []int
	for _, v1 := range nums1 {
		arrMap[v1] = true
	}
	for _, v1 := range nums2 {
		if arrMap[v1] == true {
			newNums = append(newNums, v1)
		}
	}
	return newNums
}
