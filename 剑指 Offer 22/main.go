package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	first, second := head, head
	for i := 0; i < k && first != nil; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}
	return second
}
