package main

func sortByBits(arr []int) []int {
	MyMap := map[int][]int{}
	var nums1 []int
	sort(arr)
	for i := range arr {
		k := oneBitNum(arr[i])
		MyMap[k] = append(MyMap[k], arr[i])
	}

	for i := range MyMap {
		nums1 = append(nums1, i)
	}

	sort(nums1)
	var nums2 []int
	for _, v1 := range nums1 {
		nums2 = append(nums2, MyMap[v1]...)
	}
	return nums2
}

func oneBitNum(num int) int {
	var sum int
	for num > 0 {
		sum += num % 2
		num /= 2
	}
	return sum
}
