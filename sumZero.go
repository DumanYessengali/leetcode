package main

func sumZero(n int) []int {
	arr := []int{}

	if n == 1 {
		return []int{0}
	} else if n%2 == 0 {
		for i := 1; i <= n/2; i++ {
			arr = append(arr, i)
			arr = append(arr, -1*i)
		}
	} else {
		arr = append(arr, 0)
		for i := 1; i <= (n-1)/2; i++ {
			arr = append(arr, i)
			arr = append(arr, -1*i)
		}
	}

	return arr
}
