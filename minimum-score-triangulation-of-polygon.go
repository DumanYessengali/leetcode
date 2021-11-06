package main

func minScoreTriangulation(values []int) int {
	var arr []int
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			for k := j + 1; k < len(values); k++ {
				arr = append(arr, values[i]*values[j]*values[k])
			}
		}
	}
	sort(arr)

	var sum int
	i := 0
	k := len(values) - 2
	for k > 0 {
		sum += arr[i]
		i++
		k--
	}
	return sum
}
