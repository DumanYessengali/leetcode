package main

func hammingDistance(x int, y int) int {
	diff := x ^ y
	var bits int
	for diff > 0 {
		if diff&1 != 0 {
			bits++
		}
		diff >>= 1
	}
	return bits
}
