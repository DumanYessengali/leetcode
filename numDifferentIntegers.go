package main

func numDifferentIntegers(word string) int {
	arrByte := []byte(word)

	s := ""

	var arrInt []string
	for _, v1 := range arrByte {
		if v1 > 47 && v1 < 58 {
			s += string(v1)
		} else {
			if s == "" {
				continue
			} else {
				arrInt = append(arrInt, realString(s))
			}
			s = ""
		}
	}
	if s != "" {
		arrInt = append(arrInt, realString(s))
	}
	intMap := map[string]int{}

	for _, v1 := range arrInt {
		if intMap[v1] == 1 {
			intMap[v1] = 1
		} else {
			intMap[v1] += 1

		}
	}
	return len(intMap)
}

func realString(k string) string {
	for i := 0; i < len(k); i++ {
		if k[i] == 48 {
			continue
		} else {
			return k[i:]
		}
	}
	return ""
}
