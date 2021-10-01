package main

func numUniqueEmails(emails []string) int {
	arrMap := map[string]int{}
	k := true
	m := true
	for _, v1 := range emails {
		str := ""
		for _, v2 := range v1 {
			if v2 == '@' {
				k = true
				m = false
			}
			if v2 == '+' {
				k = false
				continue
			}
			if k {
				if v2 == '.' && m {
					continue
				}
				str += string(v2)
			}
		}
		k = true
		m = true
		if arrMap[str] > 0 {
			arrMap[str] += 1
			continue
		}
		arrMap[str] = 1
	}
	return len(arrMap)
}
