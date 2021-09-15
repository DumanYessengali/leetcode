package main

func addToArrayForm(num []int, k int) []int {
	return addToFrontArray(addToArray(num, k, len(num)-1))
}

func addToArray(num []int, k, i int) ([]int, int) {
	if i >= 0 {
		num[i] += k
		k = num[i] / 10
		num[i] %= 10
		return addToArray(num, k, i-1)
	}
	return num, k
}

func addToFrontArray(num []int, k int) []int {
	if k > 0 {
		num = append([]int{k % 10}, num...)

		return addToFrontArray(num, k/10)
	}
	return num
}
