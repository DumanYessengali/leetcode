package main

func threeSum(nums []int) [][]int {
	var arr [][]int
	sort(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			l, h, sum := i+1, len(nums)-1, 0-nums[i]
			for l < h {
				if nums[l]+nums[h] == sum {
					arr = append(arr, []int{nums[i], nums[l], nums[h]})
					for l < h && nums[l] == nums[l+1] {
						l++
					}
					for l < h && nums[h] == nums[h-1] {
						h--
					}
					l++
					h--
				} else if nums[l]+nums[h] < sum {
					l++
				} else {
					h--
				}
			}
		}
	}
	return arr
}

func sort(arr []int) []int {
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
