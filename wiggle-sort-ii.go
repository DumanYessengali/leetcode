package main

// first method
func wiggleSort(nums []int) {
	sort(nums)
	n := len(nums)

	arr := make([]int, n, n)
	i := 1
	j := n - 1

	for i < n {
		arr[i] = nums[j]
		i += 2
		j--
	}
	i = 0
	for i < n {
		arr[i] = nums[j]
		j--
		i += 2
	}

	for index := range nums {
		nums[index] = arr[index]
	}
}
