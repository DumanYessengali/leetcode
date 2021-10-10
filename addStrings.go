package main

func addStrings(num1 string, num2 string) string {
	n1 := toIntArray(num1)
	n2 := toIntArray(num2)
	nums1 := addToNums(n1, n2)
	if nums1[0] == 0 {
		nums1 = nums1[1:]
	}
	return intArrayToString(nums1, 0, "")
}

func addToNums(n1, n2 []int) []int {
	if len(n1) > len(n2) {
		for len(n1) >= len(n2) {
			n2 = append([]int{0}, n2...)
		}
		n1 = append([]int{0}, n1...)
	} else {
		for len(n2) >= len(n1) {
			n1 = append([]int{0}, n1...)
		}
		n2 = append([]int{0}, n2...)

	}

	maxLen := len(n1) - 1
	maxLen2 := len(n2) - 1

	for maxLen >= 0 {
		n1[maxLen] += n2[maxLen2]
		if maxLen2-1 >= 0 {
			n2[maxLen2-1] = n1[maxLen]/10 + n2[maxLen2-1]
		}
		n1[maxLen] %= 10
		maxLen--
		maxLen2--
	}
	return n1
}

func toIntArray(arr string) []int {
	var nums []int
	for _, v1 := range arr {
		nums = append(nums, -int('0'-v1))
	}
	return nums
}

func intArrayToString(nums []int, i int, s string) string {
	if i < len(nums) {
		if nums[i] == 0 {
			s += "0"
		} else {
			s = toString(nums[i], s)
		}
		return intArrayToString(nums, i+1, s)
	}
	return s
}

func toString(num int, s string) string {
	if num > 0 {
		s = s + string(rune('0'+num%10))
		return toString(num/10, s)
	}
	if s == "" {
		return "0"
	}
	return s
}
