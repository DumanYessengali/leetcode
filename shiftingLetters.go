package main

func shiftingLetters(s string, shifts []int) string {
	var arrStr []int

	for _, v1 := range s {
		arrStr = append(arrStr, int(v1)-97)
	}

	for i := 0; i < len(shifts); i++ {
		j := 0
		for j < i+1 {
			arrStr[j] += shifts[i]
			j++
		}
		j = 0
	}

	str := ""

	for i := 0; i < len(arrStr); i++ {
		arrStr[i] %= 26
		arrStr[i] += 97
		str += string(arrStr[i])
	}

	return str
}

func shiftingLetters2(s string, shifts []int) string {
	bytes, sum := []byte(s), 0

	for i := len(s) - 1; i >= 0; i-- {
		sum += shifts[i]
		bytes[i] = byte((int(bytes[i])-'a'+sum)%26 + 'a')
	}
	return string(bytes)
}
