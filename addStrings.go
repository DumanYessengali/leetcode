package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	n1 := toIntArray(num1)
	n2 := toIntArray(num2)
	nums := addToNums(n1, n2)
	fmt.Println(nums)
	return ""
}
func addToNums(n1, n2 []int) []int {
	maxLen := len(n1) - 1
	if maxLen < len(n2)-1 {
		maxLen = len(n2) - 1
	}
	//for maxLen >= 0 {
	//
	//}
	return n1
}

func toIntArray(arr string) []int {
	var nums []int
	for _, v1 := range arr {
		nums = append(nums, -int('0'-v1))
	}
	return nums
}

func toString(num int, s string) string {
	if num > 0 {
		s = string(rune('0'+num%10)) + s
		return toString(num/10, s)
	}
	return s
}
