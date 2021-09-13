package main

func maxNumberOfBalloons(text string) int {
	byteMap := map[int32]int32{
		'b': 'b',
		'a': 'a',
		'l': 'l',
		'o': 'o',
		'n': 'n',
	}

	arrMap := map[int32]int{}

	for _, v1 := range text {
		if v1 == byteMap[v1] {
			if arrMap[v1] == 0 {
				arrMap[v1] = 1
			} else {
				arrMap[v1] += 1
			}
		}
	}
	l := arrMap['a']
	for k, v := range arrMap {
		if k == 'a' || k == 'b' || k == 'n' {
			if v < l {
				l = v
			}
		} else if k == 'l' || k == 'o' {
			if v/2 <= l && v > 1 {
				l = v / 2
			} else if v/2 > l {
				continue
			} else {
				l--
			}
		}
	}

	if l < 0 {
		return 0
	}
	return l
}
