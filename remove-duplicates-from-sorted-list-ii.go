package main

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := &ListNode{-1, head}
	prev := node
	curr := head

	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}
		curr = curr.Next
	}
	return node.Next
}
