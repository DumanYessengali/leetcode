package main

func reverseOnlyLetters(s string) string {
	arrByte := []byte(s)
	i, j := 0, len(arrByte)-1
	l, r := false, false
	for i <= j {
		if (arrByte[i] > 64 && arrByte[i] < 91) || (arrByte[i] > 96 && arrByte[i] < 123) {
			l = true
		} else {
			i++
		}
		if (arrByte[j] > 64 && arrByte[j] < 91) || (arrByte[j] > 96 && arrByte[j] < 123) {
			r = true
		} else {
			j--
		}

		if l && r {
			arrByte[i], arrByte[j] = arrByte[j], arrByte[i]
			l, r = false, false
			i++
			j--
		}
	}
	return string(arrByte)
}
