package main

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return [][]int{}
	}
	var arr [][]int
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		low, high := max(firstList[i][0], secondList[j][0]), min(firstList[i][1], secondList[j][1])

		if low <= high {
			arr = append(arr, []int{low, high})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return arr
}
