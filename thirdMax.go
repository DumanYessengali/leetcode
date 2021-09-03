package main

func thirdMax(nums []int) int {
	arr := removeDuplicate(nums)
	if len(arr) < 3 {
		max, _ := maxRemove(arr)
		return max
	} else {
		max, arr1 := maxRemove(arr)
		max, arr2 := maxRemove(arr1)
		max, _ = maxRemove(arr2)
		return max
	}
}

func removeDuplicate(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func maxRemove(nums []int) (int, []int) {
	max := nums[0]
	for _, v1 := range nums {
		if max < v1 {
			max = v1
		}
	}
	arr := []int{}
	for _, v1 := range nums {
		if max != v1 {
			arr = append(arr, v1)
		}
	}
	return max, arr
}
