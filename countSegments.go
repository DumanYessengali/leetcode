package main

func countSegments(s string) int {
	var arrStr []string
	a := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			arrStr = append(arrStr, s[a:i])
			a = i + 1
		}
	}
	k := 0
	arrStr = append(arrStr, s[a:])
	for _, v1 := range arrStr {
		if len(v1) != 0 {
			k++
		}
	}

	return k
}
