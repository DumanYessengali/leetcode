package main

func truncateSentence(s string, k int) string {
	a := 0
	for i, v1 := range s {
		if v1 == 32 {
			a++
			if a == k {
				return s[:i]
			}
		}
	}
	return s
}
