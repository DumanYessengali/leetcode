package main

func reformatNumber(number string) string {
	var arr []byte
	for i := range number {
		if number[i] > 47 && number[i] < 58 {
			arr = append(arr, number[i])
		}
	}
	if len(arr) > 3 {
		if len(arr)%3 == 1 {
			end := string(arr[len(arr)-4:])
			return toPhoneNum(string(arr[:len(arr)-4]), "") + end[0:2] + "-" + end[2:]
		} else if len(arr)%3 == 2 {
			end := string(arr[len(arr)-2:])
			return toPhoneNum(string(arr[:len(arr)-2]), "") + end[0:2]
		} else {
			str := toPhoneNum(string(arr), "")
			return str[:len(str)-1]
		}
	}
	return string(arr)
}

func toPhoneNum(str, newStr string) string {
	for i := 0; i < len(str); i += 3 {
		newStr += str[i:i+3] + "-"
	}
	return newStr
}
