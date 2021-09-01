package main

func decompressRLElist(nums []int) []int {
	var ans []int
	i := 0
	for i < len(nums) {
		j := 0
		k, l := nums[i+1], nums[i]
		for j < l {
			ans = append(ans, k)
			j++
		}
		i += 2
	}
	return ans
}
