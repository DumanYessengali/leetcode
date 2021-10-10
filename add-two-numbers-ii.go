package main

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	return reverseList(addTwoNumbers(l1, l2))
}
