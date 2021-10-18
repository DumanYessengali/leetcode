package main

func majorityElementII(nums []int) []int {
	arrMap := map[int]int{}
	var arr []int
	for i := range nums {
		arrMap[nums[i]]++

	}
	k := 0
	max := 0
	for i := range nums {
		if arrMap[nums[i]] >= k && arrMap[nums[i]] > len(nums)/3 {
			k = arrMap[nums[i]]
			max = nums[i]
			arr = checkIsUnique(arr, max)
		}
	}

	return arr
}

func checkIsUnique(arr []int, max int) []int {
	i := 0
	j := len(arr) - 1
	for i <= j {
		if arr[i] == max || arr[j] == max {
			return arr
		}
		i++
		j--
	}
	return append(arr, max)
}
