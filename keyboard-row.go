package main

func findWords(words []string) []string {
	wordMap1 := mapOfLetters("qwertyuiop")
	wordMap2 := mapOfLetters("asdfghjkl")
	wordMap3 := mapOfLetters("zxcvbnm")
	var arr []string
	for _, v1 := range words {
		n1 := 0
		n2 := 0
		n3 := 0
		for _, v2 := range v1 {
			if wordMap1[v2] == v2 {
				n1 += 1
			} else if wordMap2[v2] == v2 {
				n2 += 2
			} else if wordMap3[v2] == v2 {
				n3 += 3
			}
		}
		if len(v1) == n1 || len(v1) == n2/2 || len(v1) == n3/3 {
			arr = append(arr, v1)
		}
	}
	return arr
}

func mapOfLetters(word string) map[int32]int32 {
	wordMap := map[int32]int32{}
	for _, v1 := range word {
		wordMap[v1] = v1
		wordMap[v1-32] = v1 - 32
	}
	return wordMap
}
