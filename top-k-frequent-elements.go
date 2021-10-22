package main

func topKFrequent(nums []int, k int) []int {
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
	for i := len(arr) - 1; i > len(arr)-1-k; i-- {
		for j := range arrMap {
			if arrMap[j] == arr[i] {
				result = append(result, j)
				arrMap[j] = 0

			}
		}
	}
	return result
}
