package main

func swapPairs(head *ListNode) *ListNode {
	t := head
	if t != nil && t.Next != nil {
		t1 := t.Val
		t.Val = t.Next.Val
		t.Next.Val = t1
		t = t.Next.Next
		t = swapPairs(t)
	}
	return head
}
