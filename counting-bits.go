package main

func countBits(n int) []int {
	var arr = make([]int, n+1)

	for i := 0; i <= n; i++ {
		arr[i] = arr[i/2] + i&1
	}
	return arr
}
