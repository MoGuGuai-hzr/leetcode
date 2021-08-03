package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || right <= 1 {
		return head
	}
	newHead := reverseBetween2(head.Next, left-1, right-1)
	if left <= 1 {
		t := head.Next
		head.Next = t.Next
		t.Next = head
		return newHead
	} else {
		head.Next = newHead
		return head
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var fackHead = new(ListNode)
	fackHead.Next = head

	tail := fackHead
	for i := 0; i < left-1; i++ {
		tail = tail.Next
	}

	cur := tail.Next

	for i := left; i < right; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = tail.Next
		tail.Next = next
	}

	return fackHead.Next
}
