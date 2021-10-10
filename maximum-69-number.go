package main

func maximum69Number(num int) int {
	max := num
	k := 1
	l := 0
	m := 0
	for i := num; i > 0; i = i / 10 {
		k = k * 10
		m++
	}
	k /= 10
	for i := 0; i < m; i++ {
		if (num-l)/k != 0 {
			if (num-l)/k == 9 {
				if max < 6*k+num%k+l {
					max = 6*k + num%k + l
				}
			} else if (num-l)/k == 6 {
				if max < 9*k+num%k+l {
					max = 9*k + num%k + l
				}
			}
		}
		l = num / k * k
		k /= 10
	}
	return max
}
