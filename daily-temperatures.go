package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	arr := make([]int, n)
	day, fDay := 0, 1
	for day < n && fDay < n {
		if temperatures[day] < temperatures[fDay] {
			arr[day] = fDay - day
			day++
			fDay = day + 1
		} else {
			if fDay+1 == n {
				day++
				fDay = day + 1
			} else {
				fDay++
			}
		}

	}
	return arr
}
