package main

func reverseVowels(s string) string {
	arrMap := map[uint8]uint8{
		97:  97,
		101: 101,
		105: 105,
		111: 111,
		117: 117,
		65:  65,
		69:  69,
		73:  73,
		79:  79,
		85:  85,
	}
	arrByte := []byte(s)
	i, j := 0, len(arrByte)-1
	l, r := false, false
	for i <= j {
		if arrByte[i] == arrMap[arrByte[i]] {
			l = true
		} else {
			i++
		}

		if arrByte[j] == arrMap[arrByte[j]] {
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
