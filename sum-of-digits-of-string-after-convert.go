package main

func getLucky(s string, k int) int {
	var sum int

	for i := range s {
		for j := s[i] - 'a' + 1; j > 0; j /= 10 {
			sum += int(j % 10)
		}
	}
	k--
	var result int
	for k > 0 {
		for sum > 0 {
			result += sum % 10
			sum /= 10
		}

		sum = result
		result = 0

		k--
	}
	return sum
}
