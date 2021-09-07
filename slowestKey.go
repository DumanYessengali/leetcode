package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	arrPress := []int{}
	arrPress = append(arrPress, releaseTimes[0])

	max := arrPress[0]
	for i := 1; i < len(releaseTimes); i++ {
		arrPress = append(arrPress, releaseTimes[i]-releaseTimes[i-1])
		if max < arrPress[i] {
			max = arrPress[i]
		}
	}
	k := 0

	for _, v1 := range arrPress {
		if v1 == max {
			k++
		}
	}

	arrByte := []byte{}
	for i, v1 := range arrPress {
		if v1 == max {
			arrByte = append(arrByte, keysPressed[i])

		}
	}

	for i := len(arrByte); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arrByte[j-1] > arrByte[j] {
				arrByte[j], arrByte[j-1] = arrByte[j-1], arrByte[j]
			}
		}
	}

	return arrByte[len(arrByte)-1]
}
