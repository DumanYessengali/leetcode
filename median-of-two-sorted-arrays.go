package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	for i := len(nums1); i > 0; i-- {
		for j := 1; j < i; j++ {
			if nums1[j] < nums1[j-1] {
				nums1[j], nums1[j-1] = nums1[j-1], nums1[j]
			}
		}
	}
	if len(nums1)%2 == 0 {
		return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2-1]) / 2
	}
	return float64(nums1[len(nums1)/2])
}
