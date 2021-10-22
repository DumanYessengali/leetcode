package main

func frequencySort(s string) string {
	arrMap := map[uint8]int{}

	for i := range s {
		arrMap[s[i]]++
	}
	var arr []int
	for i := range arrMap {
		arr = append(arr, arrMap[i])
	}
	arr = sort(arr)

	var str string
	for i := len(arr) - 1; i >= 0; i-- {
		for j := range arrMap {
			if arrMap[j] == arr[i] {
				for k := 0; k < arr[i]; k++ {
					str += string(j)
				}
				arrMap[j] = 0
			}
		}
	}
	return str
}
