package main

func findDifferentBinaryString(nums []string) string {
	str := ""
	for i := 0; i < len(nums); i++ {
		if nums[i][i] == '0' {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}
