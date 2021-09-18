package main

func sumOddLengthSubarrays(arr []int) int {
	a := func(arr []int) int {
		sum := 0

		for i := 0; i < len(arr); i++ {
			s := 0
			for j := i; j < len(arr); j++ {
				s += arr[j]
				if (j-i)%2 == 0 {
					sum += s
				}
			}
		}
		return sum
	}(arr)

	return a
}
