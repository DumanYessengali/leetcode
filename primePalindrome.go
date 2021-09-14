package main

func primePalindrome(n int) int {
	if n <= 2 {
		return 2
	} else {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return primePalindrome(n + 1)
			}
		}
	}
	if !palindromeNumber(n) {
		return primePalindrome(n + 1)
	}
	return n
}

func palindromeNumber(x int) bool {
	l := 0
	if x < 0 {
		return false
	}
	for i := x; i > 0; i /= 10 {
		l = l*10 + i%10
	}
	if l == x {
		return true
	}
	return false
}
