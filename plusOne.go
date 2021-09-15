package main

func plusOne(digits []int) []int {
	return addToFrontArray(addToArray(digits, 1, len(digits)-1))
}

//func plusOne(digits []int) []int {
//	if digits[len(digits)-1]+1 == 10 {
//		if len(digits) == 1 {
//			digits[0] = 1
//			digits = append(digits, 0)
//		} else {
//			p := false
//			x := 0
//			for i := len(digits) - 1; i >= 0; i-- {
//				if digits[i]+1 == 10 {
//					x = digits[i]
//					digits[i] = 0
//
//					if i == 1 {
//						p = true
//					}
//				} else {
//					if x == 9 {
//						digits[i] += 1
//						p = false
//					}
//					break
//				}
//			}
//			if p {
//				digits[0] = 10
//			}
//			if digits[0] == 10 {
//				digits[0] = 0
//				k := digits[0]
//				m := 0
//
//				digits = append(digits, 0)
//				for i := 1; i < len(digits); i++ {
//					m = digits[i]
//					digits[i] = k
//					k = m
//				}
//				digits[0] = 1
//			}
//		}
//		return digits
//
//	}
//	digits[len(digits)-1] += 1
//	return digits
//}
