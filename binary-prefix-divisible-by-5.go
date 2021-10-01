package main

import "fmt"

func prefixesDivBy5(nums []int) []bool {
	var arrBool []bool
	fmt.Println(len(nums))
	var s string
	for i, v1 := range nums {
		s += string(rune(48 + v1))
		arrBool = append(arrBool, isDiv5(convertToIntFromBin(s)))
		fmt.Print(i, " ")
	}
	return arrBool
}

func isDiv5(n uint64) bool {
	if n%5 == 0 {
		return true
	}
	return false
}

func convertToIntFromBin(s string) uint64 {
	var n uint64
	for i := 0; i < len(s); i++ {
		k := uint64(s[len(s)-1-i]) - 48
		for j := 0; j < i; j++ {
			k *= 2
		}
		n += k
	}
	fmt.Println(n)
	return n
}
