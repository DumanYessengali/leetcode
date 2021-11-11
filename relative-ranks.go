package main

func findRelativeRanks(score []int) []string {
	var pair [][]int
	pair = append(pair, score)
	pair = append(pair, []int{})

	for i := range score {
		pair[1] = append(pair[1], i)
	}

	for i := len(pair[0]); i > 0; i-- {
		for j := 1; j < i; j++ {
			if pair[0][j] > pair[0][j-1] {
				pair[0][j], pair[0][j-1] = pair[0][j-1], pair[0][j]
				pair[1][j], pair[1][j-1] = pair[1][j-1], pair[1][j]
			}
		}
	}

	resultArr := make([]string, len(score))

	for index := range score {
		if index == 0 {
			resultArr[pair[1][index]] = "Gold Medal"
		} else if index == 1 {
			resultArr[pair[1][index]] = "Silver Medal"
		} else if index == 2 {
			resultArr[pair[1][index]] = "Bronze Medal"
		} else {
			resultArr[pair[1][index]] = reverse([]byte(toString(index+1, "")))
		}
	}

	return resultArr
}
