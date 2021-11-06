package main

func numberOfLines(widths []int, s string) []int {
	l, w := 1, 0
	for i := range s {
		width := widths[s[i]-'a']
		w += width
		if w > 100 {
			l += 1
			w = width
		}
	}
	return []int{l, w}
}
