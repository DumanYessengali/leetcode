package main

func removeElements(head *ListNode, val int) *ListNode {
	for head == nil {
		return head
	}
	current := head
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	if head.Val == val {
		return head.Next
	}
	return head
}
