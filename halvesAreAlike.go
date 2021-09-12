package main

func halvesAreAlike(s string) bool {
	r, l := []byte(s[:len(s)/2]), []byte(s[len(s)/2:])
	arrMap := map[uint8]uint8{
		97:  97,
		101: 101,
		105: 105,
		111: 111,
		117: 117,
		65:  65,
		69:  69,
		73:  73,
		79:  79,
		85:  85,
	}
	x, y := 0, 0
	for i := range r {
		if r[i] == arrMap[r[i]] {
			x++
		}
		if l[i] == arrMap[l[i]] {
			y++
		}
	}
	return y == x
}
