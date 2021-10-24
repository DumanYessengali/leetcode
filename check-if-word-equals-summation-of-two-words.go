package main

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return wordToNums(0, targetWord) == wordToNums(0, firstWord)+wordToNums(0, secondWord)
}

func wordToNums(w int, str string) int {
	for i := 0; i < len(str); i++ {
		w = w*10 + int(str[i]-'a')
	}
	return w
}
