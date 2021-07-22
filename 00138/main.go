package main

//   Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type List struct {
	new, raw *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	newHead := &Node{Val: head.Val}
	m[head] = newHead
	list := []List{{newHead, head}}
	for len(list) != 0 {
		new, raw := list[0].new, list[0].raw
		if raw.Next != nil {
			if m[raw.Next] == nil {
				new.Next = &Node{Val: raw.Next.Val}
				m[raw.Next] = new.Next
				list = append(list, List{new.Next, raw.Next})
			} else {
				new.Next = m[raw.Next]
			}
		}
		if raw.Random != nil {
			if m[raw.Random] == nil {
				new.Random = &Node{Val: raw.Random.Val}
				m[raw.Random] = new.Random
				list = append(list, List{new.Random, raw.Random})
			} else {
				new.Random = m[raw.Random]
			}
		}
		list = list[1:]
	}
	return newHead
}
