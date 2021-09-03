package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	l1, l2 := "", ""

	for _, v1 := range word1 {
		l1 += v1
	}

	for _, v2 := range word2 {
		l2 += v2
	}

	if l1 == l2 {
		return true
	}
	return false
}
