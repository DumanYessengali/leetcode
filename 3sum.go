package main

func threeSum(nums []int) [][]int {
	var arr [][]int

	if len(nums) < 2 {
		return arr
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					if nums[i] <= nums[j] && nums[j] <= nums[k] {
						arr = append(arr, []int{nums[i], nums[j], nums[k]})
					} else if nums[j] <= nums[i] && nums[i] <= nums[k] {
						arr = append(arr, []int{nums[j], nums[i], nums[k]})
					} else if nums[i] <= nums[k] && nums[k] <= nums[j] {
						arr = append(arr, []int{nums[i], nums[k], nums[j]})
					} else if nums[k] <= nums[i] && nums[i] <= nums[j] {
						arr = append(arr, []int{nums[k], nums[i], nums[j]})
					} else if nums[j] <= nums[k] && nums[k] <= nums[i] {
						arr = append(arr, []int{nums[j], nums[k], nums[i]})
					} else if nums[k] <= nums[j] && nums[j] <= nums[i] {
						arr = append(arr, []int{nums[k], nums[j], nums[i]})
					}
				}
			}
		}
	}
	return normalize(arr)
}

func arrayEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func contains(array [][]int, item []int) bool {
	for i := 0; i < len(array); i++ {
		if arrayEqual(array[i], item) {
			return true
		}
	}
	return false
}
func normalize(array [][]int) [][]int {
	var result [][]int
	for i := 0; i < len(array); i++ {
		if !contains(result, array[i]) {
			result = append(result, array[i])
		}
	}
	return result
}
