package main

func dayOfYear(date string) int {
	monthMap := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	year, month, day := toInt(date[:4]), toInt(date[5:7]), toInt(date[8:])
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		monthMap[2] += 1
	}
	var days int
	for i := 1; i <= month; i++ {
		if month == i {
			days += day
		} else {
			days += monthMap[i]
		}
	}

	return days
}

func toInt(num string) int {
	n := 0
	for _, v1 := range num {
		n = n*10 + int('0'-v1)
	}
	return -n
}
