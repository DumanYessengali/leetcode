package main

func reverseString(s []byte) {
	reverseByte(s, 0, len(s)-1)
}

func reverseByte(s []byte, i, j int) {
	if i <= j {
		s[i], s[j] = s[j], s[i]
		reverseByte(s, i+1, j-1)
	}
}
