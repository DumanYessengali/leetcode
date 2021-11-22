package main

func commonChars(words []string) []string {
	abc := make([]int, 26)

	countLetter(words[0], abc)
	for i := 1; i < len(words); i++ {
		abc2 := make([]int, 26)
		countLetter(words[i], abc2)

		for i, v1 := range abc {
			if abc2[i] < v1 {
				abc[i] = abc2[i]
			}
		}
	}
	var result []string
	for i, v1 := range abc {
		for v1 > 0 {
			v1--
			result = append(result, string('a'+i))
		}
	}
	return result
}

func countLetter(str string, abc []int) {
	for _, v1 := range str {
		abc[v1-'a']++
	}
}
