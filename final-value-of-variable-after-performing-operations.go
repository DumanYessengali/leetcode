package main

func finalValueAfterOperations(operations []string) int {
	k := 0
	for _, v1 := range operations {
		for _, v2 := range v1 {
			if v2 == '+' {
				k++
				break
			} else if v2 == '-' {
				k--
				break
			}
		}
	}
	return k
}
