package main

func sortColors(nums []int) {
	i, j := 0, len(nums)-1
	k := 0
	for i <= j {
		if nums[i] == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		} else if nums[i] == 2 {
			nums[i], nums[j] = nums[j], nums[i]
			j--
			continue
		}
		i++
	}
}
