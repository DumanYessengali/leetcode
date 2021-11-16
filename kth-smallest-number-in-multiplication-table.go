package main

func findKthNumber(m int, n int, k int) int {
	l := 1
	h := m * n

	for l < h {
		mid := l + (h-l)/2
		if !checkFunc(mid, m, n, k, 0) {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return l
}

func checkFunc(x, m, n, k, count int) bool {
	for i := 1; i <= m; i++ {
		count += min(x/i, n)
	}
	return count >= k
}
