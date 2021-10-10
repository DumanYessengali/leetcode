package main

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	for len(a) > len(b) {
		b = "0" + b
	}
	arrByteA := []byte(a)
	arrByteB := []byte(b)
	for i := len(arrByteA) - 1; i >= 0; i-- {
		if arrByteA[i] == '1' && arrByteB[i] == '1' {
			arrByteA[i] = '2'
		} else if arrByteA[i] == '1' && arrByteB[i] == '0' || arrByteA[i] == '0' && arrByteB[i] == '1' {
			arrByteA[i] = '1'
		} else {
			arrByteA[i] = '0'
		}
	}

	for i := len(arrByteA) - 1; i > 0; i-- {
		if arrByteA[i] == '2' {
			arrByteA[i] = '0'
			arrByteA[i-1] += 1
		} else if arrByteA[i] == '3' {
			arrByteA[i] = '1'
			arrByteA[i-1] += 1
		}
	}
	k := ""
	if arrByteA[0] == '2' {
		arrByteA[0] = '0'
		k = "1" + string(arrByteA)
	} else if arrByteA[0] == '3' {
		arrByteA[0] = '1'
		k = "1" + string(arrByteA)
	} else {
		k = string(arrByteA)
	}
	return k
}
