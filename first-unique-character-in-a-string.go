package main

func firstUniqChar(s string) int {
	var arr = make([]int, 26)
	for _, v1 := range s {
		arr[int(v1-'a')]++
	}

	for i, v1 := range s {
		if arr[int(v1-'a')] == 1 {
			return i
		}
	}

	return -1
}
