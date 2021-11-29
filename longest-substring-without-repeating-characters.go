package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	var l int
	var index [128]int
	for i, j := 0, 0; j < n; j++ {
		if index[s[j]] > i {
			i = index[s[j]]
		}
		if j-i+1 > l {
			l = j - i + 1
		}
		index[s[j]] = j + 1
	}
	return l
}
