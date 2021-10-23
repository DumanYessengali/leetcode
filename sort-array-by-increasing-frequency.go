package main

func frequencySortNums(nums []int) []int {
	arrMap := map[int]int{}

	for i := range nums {
		arrMap[nums[i]]++
	}
	var arr []int
	for i := range arrMap {
		arr = append(arr, arrMap[i])
	}
	arr = sort(arr)
	var result []int

	for i := 0; i < len(arr); i++ {
		for j := range arrMap {
			if arrMap[j] == arr[i] {
				for k := 0; k < arr[i]; k++ {
					result = append(result, j)

				}
				arrMap[j] = 0
			}
		}
	}
	return result
}
