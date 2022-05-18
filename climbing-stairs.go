package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	f := 1
	s := 2
	for i := 3; i <= n; i++ {
		t := f + s
		f = s
		s = t
	}
	return s
}
func climbStairsRecursion(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climb(1, 2, n)
}
func climb(f, s, n int) int {
	if n == 1 || n == 2 {
		return s
	}
	if 2 < n {
		t := f + s
		f = s
		s = t
	}
	return climb(f, s, n-1)
}
