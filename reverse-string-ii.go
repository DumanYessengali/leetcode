package main

func reverseStr(s string, k int) string {
	arrByte := []byte(s)
	l := len(arrByte)

	for i := 0; i < l; i += 2 * k {
		end := min(i+k-1, l-1)
		reverseSb(arrByte, i, end)
	}
	return string(arrByte)
}

func reverseSb(sb []byte, start, end int) {
	for start < end {
		sb[start], sb[end] = sb[end], sb[start]
		start++
		end--
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
