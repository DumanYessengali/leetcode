package main

func lengthOfLastWord(s string) int {
	var c int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			if c == 0 {
				continue
			}
			break
		}
		c++
	}
	return c
}
