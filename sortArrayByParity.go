package main

func sortArrayByParity(nums []int) []int {
	arr := []int{}

	for _, v1 := range nums {
		if v1%2 == 0 {
			arr = append(arr, v1)
		}
	}

	for _, v1 := range nums {
		if v1%2 != 0 {
			arr = append(arr, v1)
		}
	}

	return arr
}
