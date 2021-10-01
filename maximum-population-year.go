package main

func maximumPopulation(logs [][]int) int {
	arrMap := map[int]int{}
	k := false
	for _, v1 := range logs {
		for i := v1[0]; i <= v1[1]; i++ {
			if arrMap[i] > 0 {
				k = true
				arrMap[i] += 1
				continue
			}
			arrMap[i] = 1
		}
	}
	if !k {
		return logs[0][0]
	} else {
		// дорешать
		return 0
	}
}
