package main

func kWeakestRows(mat [][]int, k int) []int {
	arrMap := make([][]int, len(mat[0])+1)
	s := 0
	for i, v1 := range mat {
		for _, v2 := range v1 {
			if v2 == 0 {
				break
			}
			s++
		}
		arrMap[s] = append(arrMap[s], i)
		s = 0
	}
	var arr []int
	for j := 0; j < len(arrMap); j++ {
		if len(arrMap[j]) != 0 {
			arr = append(arr, arrMap[j]...)
		}
	}
	return arr[:k]
}
