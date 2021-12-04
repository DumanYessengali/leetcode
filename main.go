package main

func main() {
	var link *ListNode
	link = listPushBack(link, 1)
	link = listPushBack(link, 2)
	link = listPushBack(link, 3)
	link = listPushBack(link, 4)
	link = listPushBack(link, 5)

	link = removeNthFromEnd(link, 2)
	PrintList(link)
}
