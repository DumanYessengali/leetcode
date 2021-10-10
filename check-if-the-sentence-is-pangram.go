package main

func checkIfPangram(sentence string) bool {
	arrbyte := map[int32]int{}
	for _, v1 := range sentence {
		if v1 >= 96 && v1 < 123 {
			arrbyte[v1]++
		}
	}
	return len(arrbyte) == 26

}
