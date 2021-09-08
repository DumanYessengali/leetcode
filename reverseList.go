package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var link *ListNode

	if head != nil {
		for head != nil {
			link = ListPushFront(link, head.Val)
			head = head.Next
		}
	}
	return link
}

func ListPushFront(l *ListNode, data int) *ListNode {
	n := &ListNode{Val: data}
	if l == nil {
		return n
	}
	iterator := l

	fmt.Println(iterator.Val)
	iterator = n

	return iterator
}

func ListPushBack(l *ListNode, data int) *ListNode {
	n := &ListNode{Val: data}
	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func PrintList(l *ListNode) {
	it := l
	for it != nil {
		fmt.Print(it.Val, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}
