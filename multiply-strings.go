package main

func multiply(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	for len(a) > len(b) {
		b = b + "0"
	}

	//arrByteA := []byte(a)
	//arrByteB := []byte(b)
	//var arr []int
	//for i := len(arrByteA) - 1; i >= 0; i-- {
	//	k:=0
	//
	//	for j := len(arrByteB) - 1; j >= 0; j-- {
	//
	//	}
	//}
	return ""
}
