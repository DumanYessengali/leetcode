package main

func balancedStringSplit(s string) int {
	var i, j int
	for _, v1 := range s {
		if v1 == 'R' {
			i++
		} else {
			i--
		}
		if i == 0 {
			j++
		}
	}
	return j
}
