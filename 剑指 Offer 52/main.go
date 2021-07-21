package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	t := headA
	for t != nil {
		m[t] = true
		t = t.Next
	}

	t = headB
	for t != nil {
		if m[t] {
			return t
		}
		t = t.Next
	}
	return nil
}
