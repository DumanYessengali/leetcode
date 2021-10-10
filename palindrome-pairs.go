package main

func palindromePairs(words []string) [][]int {
	var arr [][]int
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if palindrome(words[i]+words[j]) && i != j {
				arr = append(arr, []int{i, j})
			}
		}
	}
	return arr
}

func palindrome(word string) bool {
	i, j := 0, len(word)-1
	for i <= j {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}
