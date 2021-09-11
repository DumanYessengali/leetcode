package main

func uniqueMorseRepresentations(words []string) int {
	morseMap := map[uint8]string{
		97:  ".-",
		98:  "-...",
		99:  "-.-.",
		100: "-..",
		101: ".",
		102: "..-.",
		103: "--.",
		104: "....",
		105: "..",
		106: ".---",
		107: "-.-",
		108: ".-..",
		109: "--",
		110: "-.",
		111: "---",
		112: ".--.",
		113: "--.-",
		114: ".-.",
		115: "...",
		116: "-",
		117: "..-",
		118: "...-",
		119: ".--",
		120: "-..-",
		121: "-.--",
		122: "--..",
	}
	var arrByte [][]byte

	for _, v1 := range words {
		arrByte = append(arrByte, []byte(v1))
	}

	morseWordMap := map[string]int{}

	for _, v1 := range arrByte {
		l := ""
		for _, v2 := range v1 {
			l += morseMap[v2]

		}
		if morseWordMap[l] == 0 {
			morseWordMap[l] = 1
		} else {
			morseWordMap[l] += 1
		}
	}

	return len(morseWordMap)
}
