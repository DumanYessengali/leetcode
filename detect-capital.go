package main

func detectCapitalUse(word string) bool {
	if checkTheCaseOfLetterTrueFalseCondition(word[0], 64, 91) {
		if len(word) == 1 {
			return true
		} else {
			if checkTheCaseOfLetterTrueFalseCondition(word[1], 96, 123) {
				return checkTheCaseOfLetter(64, 91, word)
			} else {
				return checkTheCaseOfLetter(96, 123, word)
			}
		}
	} else {
		if len(word) == 1 {
			return true
		} else {
			if checkTheCaseOfLetterTrueFalseCondition(word[1], 96, 123) {
				return checkTheCaseOfLetter(64, 91, word)
			}
		}
	}
	return false
}

func checkTheCaseOfLetter(l1, l2 uint8, word string) bool {
	for i := 2; i < len(word); i++ {
		if word[i] > l1 && word[i] < l2 {
			return false
		}
	}
	return true
}

func checkTheCaseOfLetterTrueFalseCondition(char, num1, num2 uint8) bool {
	return char > num1 && char < num2
}
