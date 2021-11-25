package main

//func selfDividingNumbers(left int, right int) []int {
//	var arr []int
//	mx := &sync.Mutex{}
//
//	for i := left; i <= right; i++ {
//		go func(i int, mx *sync.Mutex) {
//			mx.Lock()
//			defer mx.Unlock()
//			for j := i; j > 0; j /= 10 {
//				if j%10 == 0 {
//					return
//				}
//				if i%(j%10) != 0 {
//					return
//				}
//			}
//			arr = append(arr, i)
//		}(i, mx)
//	}
//	time.Sleep(time.Millisecond*1)
//	return sort(arr)
//}

func selfDividingNumbers(left int, right int) []int {
	var arr []int
	for i := left; i <= right; i++ {
		if selfDevNum(i) {
			arr = append(arr, i)
		}
	}
	return arr
}

func selfDevNum(num int) bool {
	for i := num; i > 0; i /= 10 {
		if i%10 == 0 {
			return false
		}
		if num%(i%10) != 0 {
			return false
		}
	}

	return true
}
