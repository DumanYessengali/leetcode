package main

func replaceDigits(s string) string {
	strByte := []byte(s)
	for i := 1; i < len(s); i += 2 {
		strByte[i] = strByte[i-1] + (strByte[i] - 48)
	}
	return string(strByte)
}
