package main

import (
	"fmt"
)

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
	list := new(ListNode)
	for node, carry := list, 0; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{Val: carry % 10}
		carry /= 10
	}
	return list.Next
}

func PrintList(l *ListNode) {
	it := l
	for it != nil {
		fmt.Print(it.Val, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}
