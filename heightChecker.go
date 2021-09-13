package main

func heightChecker(heights []int) int {
	arr := make([]int, len(heights))
	copy(arr, heights)

	for i := len(heights); i > 0; i-- {
		for j := 1; j < i; j++ {
			if heights[j] < heights[j-1] {
				heights[j-1], heights[j] = heights[j], heights[j-1]
			}
		}
	}
	k := 0
	for i := range heights {
		if heights[i] != arr[i] {
			k++
		}
	}

	return k
}
