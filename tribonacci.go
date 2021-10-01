package main

func tribonacci(n int) int {
	t0, t1, t2 := 0, 1, 1
	var i int

	for i = 3; i < n; i += 3 {
		t0 += t1 + t2
		t1 += t0 + t2
		t2 += t1 + t0

	}
	if i == n {
		return t0 + t1 + t2
	} else if i-1 == n {
		return t2
	} else if i-2 == n {
		return t1
	} else {
		return t0
	}
}
