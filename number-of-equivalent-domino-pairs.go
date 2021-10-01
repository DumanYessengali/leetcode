package main

func numEquivDominoPairs(dominoes [][]int) int {
	k := 0
	var a, b, c, d int
	var i int
	j := i + 1
	for i < len(dominoes)-1 {
		a = dominoes[i][0]
		b = dominoes[i][1]
		c = dominoes[j][0]
		d = dominoes[j][1]
		if j != len(dominoes) {
			if a == c && b == d || a == d && b == c {
				k++
			}
			j++
		}
		if j == len(dominoes) {
			i++

			j = i + 1
		}

	}
	return k
}

//func numEquivDominoPairs(dominoes [][]int) int {
//	k := 0
//	var a,b,c,d int
//	for i := 0; i < len(dominoes); i++ {
//		a = dominoes[i][0]
//		b= dominoes[i][1]
//		for j := i + 1; j < len(dominoes); j++ {
//			c = dominoes[j][0]
//			d = dominoes[j][1]
//			if a == c && b == d || a == d && b == c {
//				k++
//			}
//		}
//	}
//	return k
//}
