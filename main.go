package main

import "fmt"

func main() {
	var link *ListNode
	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)
	link = listPushBack(link, 7)
	link = listPushBack(link, 5)
	link = listPushBack(link, 3)
	fmt.Println(isPalindromeLinkedList(link))
}
