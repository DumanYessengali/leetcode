package main

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return s
	}
	var start, end int

	for i := range s {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		l := max(len1, len2)
		if l > end-start {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
