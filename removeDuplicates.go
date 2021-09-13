package main

func removeDuplicates(nums []int) int {
	arrMap := map[int]int{}
	for _, v1 := range nums {
		if arrMap[v1] == 0 {
			arrMap[v1] = 1
		} else {
			arrMap[v1] += 1
		}
	}
	var expectedNums []int

	for k, _ := range arrMap {
		expectedNums = append(expectedNums, k)
	}

	for i := len(expectedNums); i >= 0; i-- {
		for j := 1; j < i; j++ {
			if expectedNums[j] < expectedNums[j-1] {
				expectedNums[j], expectedNums[j-1] = expectedNums[j-1], expectedNums[j]
			}
		}
	}

	for i := 0; i < len(expectedNums); i++ {
		nums[i] = expectedNums[i]
	}
	nums = remove(nums, len(expectedNums))
	return len(nums)
}
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
