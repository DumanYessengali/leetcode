package main

import "fmt"

var symMap = map[uint8]bool{
	'!':  true,
	'?':  true,
	'\'': true,
	',':  true,
	';':  true,
	'.':  true,
}

func mostCommonWord(paragraph string, banned []string) string {
	arrByte1 := []byte(paragraph)

	var arrByte2 []byte
	for i := range arrByte1 {
		if arrByte1[i] > 64 && arrByte1[i] < 91 {
			arrByte1[i] += 32
		}
		if !symMap[arrByte1[i]] {
			arrByte2 = append(arrByte2, arrByte1[i])
			arrByte2 = append(arrByte2, ' ')
		}
	}

	str := string(arrByte2) + " "
	arrMap2 := map[string]int{}
	var a, max int
	var str2 string
	for i := 0; i < len(str); i++ {
		if str[i] == 32 {
			var k bool
			for j := range banned {
				if banned[j] == str[a:i] {
					k = true
				}
			}
			if !k {
				arrMap2[str[a:i]]++
			}
			a = i + 1
		}
	}
	fmt.Println(arrMap2)
	for i := range arrMap2 {
		if arrMap2[i] > max && i != "" {
			max = arrMap2[i]
			str2 = i
		}
	}

	return str2
}
