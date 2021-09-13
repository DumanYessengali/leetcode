package main

func insert(head *ListNode, data int) *ListNode {
	n := &ListNode{Val: data}
	if head == nil {
		return n
	} else {
		n.Next = head
		return n
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	li1 := 0
	li2 := 0
	for l1 != nil {
		li1 = li1*10 + l1.Val
		if l1.Next != nil && l1.Next.Next == nil && l1.Next.Val == 0 {
			li1 *= 10
		}
		l1 = l1.Next
	}

	for l2 != nil {
		li2 = li2*10 + l2.Val
		if l2.Next != nil && l2.Next.Next == nil && l2.Next.Val == 0 {
			li2 *= 10
		}
		l2 = l2.Next
	}
	k1 := 0
	for i := li1; i > 0; i /= 10 {
		k1 = k1*10 + i%10
	}
	k2 := 0
	for i := li2; i > 0; i /= 10 {
		k2 = k2*10 + i%10
	}
	list3 := k1 + k2

	var newList *ListNode
	var list4 []int
	for i := list3; i > 0; i /= 10 {
		list4 = append(list4, i%10)
	}
	if list3 == 0 {
		newList = insert(newList, 0)
	}
	for i := len(list4) - 1; i >= 0; i-- {
		newList = insert(newList, list4[i])
	}
	return newList
}
