package main

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	map1 := toMap(nums1, map[int]int{})
	map2 := toMap(nums2, map[int]int{})
	map3 := toMap(nums3, map[int]int{})

	map1 = unionTwoMap(map1, map2)
	map1 = unionTwoMap(map1, map3)
	var arr []int

	for k := range map1 {
		if map1[k] > 1 {
			arr = append(arr, k)
		}
	}

	return arr
}

func toMap(nums []int, map1 map[int]int) map[int]int {
	for i := range nums {
		map1[nums[i]] = 1
	}
	return map1
}

func unionTwoMap(map1 map[int]int, map2 map[int]int) map[int]int {
	for k := range map2 {
		map1[k] += map2[k]
	}
	return map1
}
