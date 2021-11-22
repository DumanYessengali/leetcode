package main

func mergeIntervals(in [][]int) [][]int {
	if len(in) <= 1 {
		return in
	}
	for i := len(in); i > 0; i-- {
		for j := 1; j < i; j++ {
			if in[j][0] < in[j-1][0] {
				in[j], in[j-1] = in[j-1], in[j]
			}
		}
	}
	arr := append([][]int{}, in[0])
	for i := 1; i < len(in); i++ {
		l := len(arr) - 1
		if arr[l][1] >= in[i][0] && (arr[l][0] <= in[i][0] || arr[l][1] <= in[i][1] || arr[l][0] <= in[i][1]) {
			arr[l][1] = max(in[i][1], arr[l][1])
			arr[l][0] = min(in[i][0], arr[l][0])
		} else {
			arr = append(arr, in[i])
			if arr[l][1] < arr[len(arr)-2][0] {
				arr[l], arr[len(arr)-2] = arr[len(arr)-2], arr[l]
			}
		}
	}
	return arr
}
