package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	lenth := 0
	for cur := head; cur != nil; lenth++ {
		cur = cur.Next
	}

	lenth -= lenth % k

	newHead := new(ListNode)
	prev := newHead
	prev.Next = head
	cur := prev.Next

	for i := 1; cur != nil && i < lenth; i++ {
		if i%k == 0 {
			prev = cur
			cur = prev.Next
			continue
		}

		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return newHead.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	lenth := 0
	for cur := head; cur != nil; lenth++ {
		cur = cur.Next
	}

	lenth -= lenth % k

	newHead := new(ListNode)
	prev := newHead
	prev.Next = head

	for i := 1; i < lenth; i += k {
		prev = reverse(prev, k)
	}

	return newHead.Next
}

func reverse(prev *ListNode, lenth int) (tail *ListNode) {
	cur := prev.Next
	for i := 1; i < lenth; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return cur
}
