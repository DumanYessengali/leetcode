package main

func nextGreaterElements(nums []int) []int {
	var arr = make([]int, len(nums), cap(nums))
	for i := range nums {
		arr[i] = -1
		myArr := append(nums[i:], nums[:i]...)
		for k := range myArr {
			if myArr[k] > nums[i] {
				arr[i] = myArr[k]
				break
			}
		}
	}
	return arr
}
