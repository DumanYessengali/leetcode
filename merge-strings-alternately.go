package main

func mergeAlternately(word1 string, word2 string) string {
	min := word2
	if len(word1) < len(word2) {
		min = word1
	}
	var arr []byte
	for i := 0; i < len(min); i++ {
		arr = append(arr, word1[i])
		arr = append(arr, word2[i])
	}
	if word1 == min {
		return string(arr) + word2[len(word1):]
	} else {
		return string(arr) + word2[len(word2):]
	}
}
