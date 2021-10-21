package main

func trimMean(arr []int) float64 {
	arr = sort(arr)

	var sum float64
	for i := len(arr) / 20; i < len(arr)-(len(arr)/20); i++ {

		sum += float64(arr[i])
	}
	return sum / float64(len(arr)-(len(arr)/20)*2)
}
