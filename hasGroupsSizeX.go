package main

func hasGroupsSizeX(deck []int) bool {
	arrMap := map[int]int{}
	for _, v1 := range deck {
		if arrMap[v1] == 0 {
			arrMap[v1] = 1
		} else {
			arrMap[v1] += 1
		}
	}
	max := arrMap[deck[0]]
	for _, v1 := range arrMap {
		if max < v1 {
			max = v1
		}
	}

	if max < 2 {
		return false
	}

	for _, v1 := range arrMap {
		if max%v1 != 0 {
			return false
		}
	}
	return true
}
