package main

func reversePrefix(word string, ch byte) string {
	for i, v1 := range word {
		if v1 == int32(ch) {
			return reverse([]byte(word[:i])) + word[i:]
		}
	}
	return word
}
