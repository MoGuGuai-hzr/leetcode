package main

import (
	"container/heap"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists1(lists []*ListNode) *ListNode {
	var merge func(list1, list2 *ListNode) *ListNode

	merge = func(list1, list2 *ListNode) *ListNode {
		if list1 == nil {
			return list2
		}
		if list2 == nil {
			return list1
		}
		var head *ListNode
		if list1.Val >= list2.Val {
			head = list2
			head.Next = merge(list1, list2.Next)
		} else {
			head = list1
			head.Next = merge(list1.Next, list2)
		}
		return head
	}
	var rst *ListNode
	for _, list := range lists {
		rst = merge(rst, list)
	}
	return rst
}

type hint []int

func (h *hint) Len() int {
	return len(*h)
}

func (h *hint) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Swap swaps the elements with indexes i and j.
func (h *hint) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hint) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *hint) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func mergeKLists2(lists []*ListNode) *ListNode {
	h := hint{}
	for _, list := range lists {
		for list != nil {
			heap.Push(&h, list.Val)
			list = list.Next
		}
	}
	var head = new(ListNode)
	t := head
	for len(h) != 0 {
		t.Next = new(ListNode)
		t = t.Next
		t.Val = heap.Pop(&h).(int)
	}
	return head.Next
}

func mergeKLists3(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	middle := len(lists) / 2
	l1 := mergeKLists3(lists[:middle])
	l2 := mergeKLists3(lists[middle:])

	var merge func(list1, list2 *ListNode) *ListNode

	merge = func(list1, list2 *ListNode) *ListNode {
		if list1 == nil {
			return list2
		}
		if list2 == nil {
			return list1
		}
		var head *ListNode
		if list1.Val >= list2.Val {
			head = list2
			head.Next = merge(list1, list2.Next)
		} else {
			head = list1
			head.Next = merge(list1.Next, list2)
		}
		return head
	}

	return merge(l1, l2)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	middle := len(lists) / 2
	l1 := mergeKLists(lists[:middle])
	l2 := mergeKLists(lists[middle:])

	var merge func(list1, list2 *ListNode) *ListNode

	merge = func(list1, list2 *ListNode) *ListNode {
		if list1 == nil {
			return list2
		}
		if list2 == nil {
			return list1
		}
		var head, cur *ListNode
		prev1 := list1
		prev2 := list2
		if prev1.Val <= prev2.Val {
			head = prev1
			cur = head
			prev1 = prev1.Next
		} else {
			head = prev2
			cur = head
			prev2 = prev2.Next
		}

		for prev1 != nil && prev2 != nil {
			if prev1.Val <= prev2.Val {
				cur.Next = prev1
				cur = prev1
				prev1 = prev1.Next
			} else {
				cur.Next = prev2
				cur = prev2
				prev2 = prev2.Next
			}
		}

		if prev1 != nil {
			cur.Next = prev1
		} else {
			cur.Next = prev2
		}

		return head
	}

	return merge(l1, l2)
}
