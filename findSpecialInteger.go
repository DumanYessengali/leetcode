package main

func findSpecialInteger(arr []int) int {
	arrMap := map[int]int{}

	for _, v1 := range arr {
		if arrMap[v1] == 0 {
			arrMap[v1] = 1
		} else {
			arrMap[v1] += 1
		}
	}

	max := 0
	for i, v1 := range arrMap {
		if float64(v1)/float64(len(arr))*100 > 25.0 {
			max = i
		}
	}

	return max
}
