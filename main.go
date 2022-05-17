package main

func main() {
	var list *ListNode
	list = listPushBack(list, 1)
	list = listPushBack(list, 2)
	//list = listPushBack(list, 3)
	//list = listPushBack(list, 4)

	PrintList(swapPairs(list))
}
