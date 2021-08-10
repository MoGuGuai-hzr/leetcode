package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	if n <= 2 {
		return
	}

	middle := n / 2
	prev := head
	for i := 0; i < middle; i++ {
		prev = prev.Next
	}

	cur = prev.Next
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	cur = head
	for cur.Next != nil && prev.Next != nil {
		next := prev.Next
		prev.Next = next.Next
		next.Next = cur.Next
		cur.Next = next
		cur = next.Next
	}
}
