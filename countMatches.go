package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	n := 0
	if ruleKey == "type" {
		n = 1
	} else if ruleKey == "color" {
		n = 2
	} else {
		n = 3
	}

	l := 0

	for _, v1 := range items {
		for i, v2 := range v1 {
			if i+1 == n && v2 == ruleValue {
				l++
			}
		}
	}
	return l
}
