package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	k := 0
	for i := m; i < n+m; i++ {
		if nums1[i] == 0 {
			nums1[i] = nums2[k]
			k++
		}
	}

	for i := len(nums1); i > 0; i-- {
		for j := 1; j < i; j++ {
			if nums1[j-1] > nums1[j] {
				nums1[j], nums1[j-1] = nums1[j-1], nums1[j]
			}
		}
	}
}
