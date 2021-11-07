package main

func multiply(a string, b string) string {
	length := len(a) + len(b)
	a, b = reverse([]byte(a)), reverse([]byte(b))
	reversedResult := make([]int32, length)
	for index2, v2 := range b {
		for index1, v1 := range a {
			currentPosition := index1 + index2
			carry := reversedResult[currentPosition]
			var multiplication int32
			if carry == 0 {
				multiplication = (v1-'0')*(v2-'0') + carry
			} else {
				multiplication = (v1-'0')*(v2-'0') + carry - '0'
			}
			reversedResult[currentPosition] = multiplication%10 + '0'
			reversedResult[currentPosition+1] += multiplication/10 + '0'
			if reversedResult[currentPosition] > 57 {
				reversedResult[currentPosition] -= '0'
			}
			if reversedResult[currentPosition+1] > 57 {
				reversedResult[currentPosition+1] -= '0'
			}
		}
	}
	result := reverse([]byte(string(reversedResult)))
	for index := range result {
		if result[index] > '0' {
			return result[index:]
		}
	}
	return "0"
}
