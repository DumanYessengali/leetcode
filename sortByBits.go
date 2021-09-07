package main

import (
	"fmt"
	"math"
)

func changeToBits(num int) string {
	str := ""

	for num > 0 {
		if num%2 == 0 {
			str += "0"
		} else {
			str += "1"
		}
		num /= 2
	}

	realStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		realStr += string(str[i])
	}

	return realStr
}

func sortByBits(arr []int) []int {
	arrMap := map[string]int{}
	mapRep := map[int]int{}

	for _, v1 := range arr {
		n := changeToBits(v1)

		arrMap[n] = func(n string) int {
			k := 0
			for _, v1 := range n {
				if v1 == 49 {
					k++
				}
			}
			return k
		}(n)

		if mapRep[arrMap[n]] == 0 {
			mapRep[arrMap[n]] = 1
		} else {
			mapRep[arrMap[n]] += 1
		}
	}
	fmt.Println(mapRep)
	fmt.Println(arrMap)
	newArr := make([][]int, len(mapRep))

	for i, v1 := range mapRep {
		for j := 0; j < v1; j++ {

			for k, v3 := range arrMap {

				if v3 == i {
					newArr[i][j] = convertToInt(k)
				}
			}
		}
	}
	fmt.Println(newArr)
	return []int{}
}

func convertToInt(k string) int {
	sum := 0
	for i := 0; i < len(k); i++ {
		sum = sum + (int(k[i])-48)*int(math.Pow(2.0, float64(i)))
	}
	return sum
}
