package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if n > 0 {
			if flowerbed[i] == 0 &&
				(i == 0 || flowerbed[i-1] == 0) &&
				(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
				n -= 1
				i++
			}
		} else {
			break
		}
	}

	return n == 0
}
