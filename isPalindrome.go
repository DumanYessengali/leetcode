package main

func isPalindrome(s string) bool {
	var strArr []int
	for _, v1 := range s {
		if v1 > 64 && v1 < 91 {
			strArr = append(strArr, int(v1)+32)
		} else if v1 > 96 && v1 < 123 {
			strArr = append(strArr, int(v1))
		} else if v1 > 47 && v1 < 58 {
			strArr = append(strArr, int(v1))
		}
	}

	i := 0
	j := len(strArr) - 1

	for i <= j {
		if strArr[i] != strArr[j] {
			return false
		}
		i++
		j--
	}

	return true
}
