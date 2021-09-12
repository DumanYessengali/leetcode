package main

func reverseWords1(s string) string {
	var arrStr []string
	a := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			arrStr = append(arrStr, s[a:i])
			a = i + 1
		}
	}
	arrStr = append(arrStr, s[a:])
	return reverseStrings(arrStr, 0, len(arrStr)-1)
}

func reverseStrings(arrStr []string, i, j int) string {
	c1, c2 := false, false

	for i <= j {
		if arrStr[i] != "" {
			c1 = true
		} else {
			i++
		}

		if arrStr[j] != "" {
			c2 = true
		} else {
			j--
		}

		if c1 && c2 {
			arrStr[i], arrStr[j] = arrStr[j], arrStr[i]
			i++
			j--
			c1, c2 = false, false
		}

	}
	k := ""

	for _, v1 := range arrStr {
		if v1 != "" {
			k += v1 + " "
		}
	}
	return k[:len(k)-1]
}
