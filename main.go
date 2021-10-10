package main

func main() {
	var link *ListNode
	link = insert(link, 1)
	link = insert(link, 1)
	PrintList(deleteDuplicates(link))
}
