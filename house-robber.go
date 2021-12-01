package main

func rob(nums []int) int {
	var pMax, cMax int
	for i := range nums {
		t := cMax
		if pMax+nums[i] > cMax {
			cMax = pMax + nums[i]
		}
		pMax = t
	}
	return cMax
}
