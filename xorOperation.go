package main

func xorOperation(n int, start int) int {
	nums := 0
	for i := 0; i < n; i++ {
		nums ^= start + 2*i
	}
	return nums
}
