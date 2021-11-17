package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var arr = make([]int, len(nums1), cap(nums1))

	for i := range nums1 {
		arr[i] = -1
	LOOP:
		for j := range nums2 {
			if nums1[i] == nums2[j] && j+1 < len(nums2) {
				for k := range nums2[j:] {
					if nums2[k+j] > nums2[j] {
						arr[i] = nums2[k+j]
						break LOOP
					}
				}
			}
		}
	}
	return arr
}
