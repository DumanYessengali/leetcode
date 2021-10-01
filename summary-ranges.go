package main

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	var str []string

	f := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 != nums[i] {
			sum := intToString(f)

			if nums[i-1] != f {
				sum += "->" + intToString(nums[i-1])
			}
			f = nums[i]
			str = append(str, sum)
		}
	}
	s := intToString(f)
	if nums[len(nums)-1] != f {
		s += "->" + intToString(nums[len(nums)-1])
	}
	str = append(str, s)

	return str
}

func intToString(num int) string {
	var s string
	var tf bool
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num *= -1
		tf = true
	}
	var nums []string
	for num > 0 {
		nums = append(nums, string(rune('0'+num%10)))
		num /= 10
	}

	i, j := 0, len(nums)-1
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	for _, v1 := range nums {
		s += v1
	}

	if tf {
		return "-" + s
	}
	return s
}
