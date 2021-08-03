package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = newHead
		newHead = curr
		curr = next
	}
	return newHead
}
