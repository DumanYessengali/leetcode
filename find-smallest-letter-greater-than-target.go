package main

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, v1 := range letters {
		if v1 > target {
			return v1
		}
	}
	return letters[0]
}
