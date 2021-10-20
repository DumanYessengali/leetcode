package main

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	i := 1
	k := n
	for n > 0 {
		if 2*i <= k {
			arr[2*i] = arr[i]
		}
		if 2*i+1 <= k {
			arr[2*i+1] = arr[i] + arr[i+1]
		}
		i++
		n--
	}
	max := arr[0]

	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
