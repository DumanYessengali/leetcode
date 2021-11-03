package main

func isAnagram(s string, t string) bool {
	s1 := toStringMap(s, map[uint8]int{})
	s2 := toStringMap(t, map[uint8]int{})

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return len(s1) == len(s2)
}

func toStringMap(s string, strMap map[uint8]int) map[uint8]int {
	for i := range s {
		strMap[s[i]]++
	}
	return strMap
}
