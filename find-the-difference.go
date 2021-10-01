package main

func findTheDifference(s string, t string) byte {
	arrMap := map[int32]int{}

	for _, v1 := range t {
		arrMap[v1]++
	}
	for _, v1 := range s {
		arrMap[v1]--
	}

	for i, v1 := range arrMap {
		if v1 != 0 {
			return byte(i)
		}
	}
	return 'a'
}
