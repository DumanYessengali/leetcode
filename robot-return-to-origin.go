package main

func judgeCircle(moves string) bool {
	var x, y int
	for i := range moves {
		if moves[i] == 'L' {
			x++
		} else if moves[i] == 'R' {
			x--
		} else if moves[i] == 'U' {
			y++
		} else if moves[i] == 'D' {
			y--
		}
	}
	return x == 0 && y == 0
}
