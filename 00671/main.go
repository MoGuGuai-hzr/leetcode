package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	rst := -1
	if root == nil {
		return rst
	}
	stack := []*TreeNode{root}
	min := root.Val
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			// 小的值入栈
			if node.Left.Val > min {
				if rst == -1 || node.Left.Val < rst {
					rst = node.Left.Val
				}
				stack = append(stack, node.Right)
			} else if node.Right.Val > min {
				if rst == -1 || node.Right.Val < rst {
					rst = node.Right.Val
				}
				stack = append(stack, node.Left)
			} else {
				x := findSecondMinimumValue(node.Left)
				if rst == -1 || x != -1 && rst > x {
					rst = x
				}
				y := findSecondMinimumValue(node.Right)
				if rst == -1 || y != -1 && rst > y {
					rst = y
				}
			}
		}
	}
	return rst
}
