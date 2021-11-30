package main

func convert(s string, n int) string {
	strLen := len(s)
	var m, colBetween, count int
	if n-2 > 0 {
		colBetween = n - 2
	}

	for i := strLen; i > 0; {
		m++
		if i-n > 0 {
			i -= n
		} else {
			break
		}

		if i-colBetween > 0 {
			m += colBetween
			i -= colBetween
		} else {
			m += i
			i = 0
		}
	}

	matrix := make([][]string, n)
	for i := range matrix {
		matrix[i] = make([]string, m)
	}
	var i, j int
	if n == 1 {
		for i, v1 := range matrix {
			for j := range v1 {
				matrix[i][j] = string(s[count])
				count++
			}
		}
	} else {
		for count != strLen {
			if i < n {
				iToN(&i, &j, &count, matrix, s)
			} else {
				i -= 1
				iToZero(&i, &j, &count, matrix, s)
			}
		}
	}
	var str string
	for _, v1 := range matrix {
		for _, v2 := range v1 {
			if v2 != "" {
				str += v2
			}
		}
	}
	return str
}

func iToZero(i, j, count *int, matrix [][]string, s string) {
	strL := len(s)

	for *i > 0 {
		if *count == strL {
			break
		}
		*i--
		*j++
		matrix[*i][*j] = string(s[*count])
		*count++

	}
}
func iToN(i, j, count *int, matrix [][]string, s string) {
	if matrix[*i][*j] != "" {
		*i++
	}
	strL := len(s)

	for *i < len(matrix) {
		if *count == strL {
			break
		}
		matrix[*i][*j] = string(s[*count])
		*count++
		*i++
	}
}
