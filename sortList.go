package main

import "fmt"

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := middleOfList(head)

	newtOfMiddle := middle.Next
	middle.Next = nil
	left := sortList(head)
	right := sortList(newtOfMiddle)

	sortedList := mergeSort(left, right)
	return sortedList
}

func mergeSort(left, right *ListNode) *ListNode {
	var result *ListNode
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Val <= right.Val {
		result = left
		result.Next = mergeSort(left.Next, right)
	} else {
		result = right
		result.Next = mergeSort(left, right.Next)
	}
	return result
}

func middleOfList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func Print(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
