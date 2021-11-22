package main

import "fmt"

func main() {
	fmt.Println(mergeIntervals([][]int{{1, 4}, {4, 5}}))
	fmt.Println(mergeIntervals([][]int{{1, 4}, {0, 4}}))
	fmt.Println(mergeIntervals([][]int{{1, 4}, {0, 0}}))
	fmt.Println(mergeIntervals([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}
