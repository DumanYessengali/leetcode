package main

func kthDistinct(arr []string, k int) string {
	var arrStr []string
	for i := range arr {
		var k bool
		for j := range arr {
			if i != j && arr[j] == arr[i] {
				k = true
			}
		}
		if !k {
			arrStr = append(arrStr, arr[i])
		}
	}
	if len(arrStr) < k {
		return ""
	}
	return arrStr[k-1]
}
