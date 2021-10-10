package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = deleteDuplicates(head.Next.Next)
			continue
		}
		head = head.Next
	}
	return node
}
