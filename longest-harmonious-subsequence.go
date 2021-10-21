package main

func findLHS(nums []int) int {
	k := 0
	for i := range nums {
		var name int
		var t bool
		for j := range nums {
			if nums[i] == nums[j] {
				name++
			} else if nums[i] == nums[j]+1 {
				t = true
				name++
			}
		}
		if t {
			if name > k {
				k = name
			}
		}
	}
	return k
}
