package main

func maxPower(s string) int {
	consec := 1

	k := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			k++
		} else {
			consec = max(consec, k)
			k = 0
		}
	}

	return max(consec, k)
}
