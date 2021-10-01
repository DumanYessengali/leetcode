package main

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	max := 0
	max2 := 0
	j := 0
	for i, v1 := range nums {
		if max < v1 {
			max = v1
			j = i
		}
	}
	for _, v1 := range nums {
		if v1 != max {
			if v1 > max2 {
				max2 = v1
			}
		}
	}
	if max >= max2*2 {
		return j
	}
	return -1
}
