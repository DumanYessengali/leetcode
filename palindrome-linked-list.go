package main

func isPalindromeLinkedList(head *ListNode) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var prev *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}
	left, right := head, prev

	for right != nil {
		if right.Val != left.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
