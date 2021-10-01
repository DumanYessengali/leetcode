package main

//func findMaxConsecutiveOnes(nums []int) int {
//	max := 0
//	realMax := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == 1 {
//			max++
//		} else {
//			if realMax < max {
//				realMax = max
//			}
//			max = 0
//		}
//	}
//	if realMax < max {
//		return max
//	}
//	return realMax
//}
func findMaxConsecutiveOnes(nums []int) int {
	var s, m int
	for _, v1 := range nums {
		if v1 == 1 {
			m++
		} else {
			m = 0
		}
		if m > s {
			s = m
		}
	}
	return s
}
