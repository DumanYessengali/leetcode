package main

func areOccurrencesEqual(s string) bool {
	m := map[uint8]int{}
	for i := 0; i < len(s); i++ {
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
