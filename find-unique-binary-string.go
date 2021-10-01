package main

func findDifferentBinaryString(nums []string) string {
	maxOne := 0
	//maxStr := ""
	for _, v1 := range nums {
		k := 0
		for _, v2 := range v1 {
			if v2 == '1' {
				k++
			}
		}
		if maxOne < k {
			maxOne = k
		}
	}
	return ""
}
