package main

func getDecimalValue(head *ListNode) int {
	number := head.Val
	for head.Next != nil {
		number = number*2 + head.Next.Val
		head = head.Next
	}
	return number
}
