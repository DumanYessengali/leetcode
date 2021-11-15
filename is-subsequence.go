package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	var i, j int

	for j < len(t) {
		if s[i] == t[j] {
			i++
			if i == len(s) {
				return true
			}
		}
		j++
	}
	return false
}
