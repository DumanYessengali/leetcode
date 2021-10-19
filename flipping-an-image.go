package main

func flipAndInvertImage(image [][]int) [][]int {
	for i := range image {
		x, y := 0, len(image[i])-1
		for x <= y {
			image[i][x], image[i][y] = image[i][y], image[i][x]
			if image[i][x] == 1 {
				image[i][x] = 0
			} else {
				image[i][x] = 1
			}
			if image[i][y] == 1 {
				image[i][y] = 0
			} else {
				image[i][y] = 1
			}
			x++
			y--
		}
	}
	return image
}
