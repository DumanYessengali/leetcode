package main

func letterCombinations(digits string) []string {
	arrMap := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	comb := 0
	m := 1
	for _, d := range digits {
		m *= len(arrMap[d])
	}
	if m > 2 {
		comb = m
	}
	res := make([]string, comb)
	for _, d := range digits {
		comb = comb / len(arrMap[d])
		j := 0
		for j < len(res) {
			for _, v := range arrMap[d] {
				for i := 0; i < comb; i++ {
					res[j] += v
					j++
				}
			}

		}
	}
	return res
}
