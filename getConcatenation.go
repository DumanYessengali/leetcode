package main

func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		if 0 <= i && i < len(nums)*2 {
			ans[i] = nums[i]
			ans[i+len(nums)] = nums[i]
		}
	}
	return ans

}
