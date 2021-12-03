package main

func maxProduct(nums []int) int {
	minV, maxV, maxP := nums[0], nums[0], nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			minV, maxV = maxV, minV
		}
		minV = min(nums[i], nums[i]*minV)
		maxV = max(nums[i], nums[i]*maxV)
		if maxV > maxP {
			maxP = maxV
		}
	}
	return maxP
}

func mergeS(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merges(mergeS(left), mergeS(right))
}

func merges(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
