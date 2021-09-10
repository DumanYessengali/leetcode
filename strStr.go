package main

func strStr(haystack string, needle string) int {
	if haystack == needle || len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if len(haystack[i:]) >= len(needle) && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
