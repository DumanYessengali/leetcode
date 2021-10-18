package main

func containsNearbyDuplicate(nums []int, k int) bool {
	arrMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		idx, ok := arrMap[nums[i]]
		if ok && i-idx <= k {
			return true
		}
		arrMap[nums[i]] = i
	}

	return false
}
