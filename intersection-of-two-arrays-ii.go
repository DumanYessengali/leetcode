package main

func intersect(nums1 []int, nums2 []int) []int {
	arrMap := map[int]int{}
	var arr []int
	for _, v1 := range nums1 {
		arrMap[v1] += 1
	}

	for _, v1 := range nums2 {
		if arrMap[v1] > 0 {
			arrMap[v1] -= 1
			arr = append(arr, v1)
		}
	}

	return arr
}
