package main

//func commonChars(words []string) []string {
//	for i, v1 := range words {
//		words[i] = sortByBytes([]byte(v1))
//	}
//	var letters []string
//	word := words[0]
//
//	tf := false
//	for i := 0; i < len(word)-1; i++ {
//		k:=0
//		for j := 1; j < len(words)-1; j++ {
//			if wor
//		}
//	}
//	return letters
//}

func sortByBytes(arrB []byte) string {
	for i := len(arrB); i > 0; i-- {
		for j := 1; j < len(arrB); j++ {
			if arrB[j-1] > arrB[j] {
				arrB[j], arrB[j-1] = arrB[j-1], arrB[j]
			}
		}
	}
	return string(arrB)
}
