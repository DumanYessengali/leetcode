package main

func rotateString(s string, goal string) bool {
	arrByte := []byte(s)

	for i := 0; i < len(arrByte); i++ {
		for j := 1; j < len(arrByte); j++ {
			arrByte[j], arrByte[j-1] = arrByte[j-1], arrByte[j]
		}
		if string(arrByte) == goal {
			return true
		}
	}
	return false
}
