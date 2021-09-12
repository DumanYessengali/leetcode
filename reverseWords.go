package main

func reverseWords(s string) string {
	var a int
	var str string

	for i, v1 := range s {
		if v1 == 32 {
			str += reverse([]byte(s[a:i])) + " "
			a = i + 1
		} else if len(s)-1 == i {
			str += reverse([]byte(s[a:]))
		}
	}

	return str
}

func reverse(word []byte) string {
	i := 0
	j := len(word) - 1
	for i <= j {
		word[i], word[j] = word[j], word[i]
		i++
		j--
	}
	return string(word)
}
