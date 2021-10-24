package main

func areOccurrencesEqual(s string) bool {
	m := map[uint8]int{}
	for i := range s {
		m[s[i]]++
	}
	k := m[s[0]]
	for i := range m {
		if k != m[i] {
			return false
		}
	}
	return true
}
