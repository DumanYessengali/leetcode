package main

func maxRepeating(sequence string, word string) int {
	if len(word) == 0 {
		return 0
	} else if sequence == word {
		return 1
	}
	k := 0
	for i := 0; i < len(sequence); i++ {
		if len(sequence[i:]) >= len(word) && sequence[i:i+len(word)] == word {
			k++
			i += len(word) - 1
		}
	}
	return k
}
