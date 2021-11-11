package main

func minStartValue(nums []int) int {
	firstNum, firstBool := 1, true
	var i int
	current := firstNum
	for firstBool {
		if len(nums)-1 == i && firstNum+nums[i] > 0 {
			break
		}

		if firstNum+nums[i] < 1 {
			current++
			firstNum = current
			i = 0
			continue
		}

		firstNum += nums[i]
		i++
	}
	return current
}
