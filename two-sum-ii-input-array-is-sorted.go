package main

func twoSum(numbers []int, target int) []int {
	return findTarget(numbers, target, numbers[0]+numbers[len(numbers)-1], 0, len(numbers)-1)
}

func findTarget(numbers []int, target, sum, i, j int) []int {
	if sum != target {
		if sum < target {
			i++
		} else {
			j--
		}

		sum = numbers[i] + numbers[j]
		return findTarget(numbers, target, sum, i, j)
	}
	return []int{i + 1, j + 1}
}
