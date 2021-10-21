package main

func maximumPopulation(logs [][]int) int {
	arrMap := map[int]int{}
	var max int

	for _, v1 := range logs {
		for i := v1[0]; i < v1[1]; i++ {
			arrMap[i]++
			if max < arrMap[i] {
				max = arrMap[i]
			}
		}
	}
	var arr []int
	for i := range arrMap {
		if arrMap[i] == max {
			arr = append(arr, i)
		}
	}

	return sort(arr)[0]
}
