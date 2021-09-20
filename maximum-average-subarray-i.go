package main

func findMaxAverage(nums []int, k int) float64 {
	var sumMax int
	if len(nums) == k {
		sumMax = sum(nums)
	} else {
		sumMax = sum(nums[0:k])
		for i := 0; i < len(nums)-k+1; i++ {
			m := sum(nums[i : i+k])
			if m > sumMax {
				sumMax = m
			}
		}
	}
	//fmt.Println(sumMax/k, float64(sumMax-(sumMax/k)*k)/float64(k))
	return float64(sumMax/k) + float64(sumMax-(sumMax/k)*k)/float64(k)
}

func sum(nums []int) int {
	var sum int
	for _, v1 := range nums {
		sum += v1
	}
	return sum
}
