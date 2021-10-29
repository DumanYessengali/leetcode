package main

import "fmt"

func main() {
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82))
}
