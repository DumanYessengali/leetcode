package main

func maxLengthBetweenEqualCharacters(s string) int {
	var max int
	var k bool
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] && i != j {
				if max < j-i-1 {
					max = j - i - 1
				}
				k = true
			}
		}
	}
	if k {
		return max
	}
	return -1
}
