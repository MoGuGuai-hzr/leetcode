package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var n int
	for t := head; t != nil; t = t.Next {
		n++
	}
	cnt := n / k
	for i := 0; i < cnt; i++ {
		h := head
		p := head
		for j := 0; j < k; j++ {
			h.Next = h.Next.Next
		}
	}
	return nil
}
