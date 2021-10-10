package main

func largestTriangleArea(points [][]int) float64 {
	var m float64
	for _, i := range points {
		for _, j := range points {
			for _, k := range points {
				z := 0.5 * float64(i[0]*j[1]+j[0]*k[1]+k[0]*i[1]-j[0]*i[1]-k[0]*j[1]-i[0]*k[1])
				if z < 0 {
					if m < (-1 * z) {
						m = z * -1
					}
				} else {
					if m < z {
						m = z
					}
				}
			}
		}
	}
	return m
}
