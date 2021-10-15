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
