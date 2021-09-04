package main

func countPrimes(n int) int {
	a := 0

	if n <= 2 {
		return a
	}

	prime := make([]bool, n+1)

	for i := 0; i <= n; i++ {
		prime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if prime[i] == true {
			for j := i * i; j <= n; j += i {
				prime[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if prime[i] == true {
			a++
		}
	}

	if prime[n] == true {
		return a - 1
	}

	return a
}
