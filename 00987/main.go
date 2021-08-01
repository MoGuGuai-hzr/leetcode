package main

import (
	"math"
	"sort"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Elem struct {
	row, val int
}

type List1 []*Elem

func (list *List1) Insert(elem *Elem) {
	l := *list
	n := len(l)
	i := sort.Search(n, func(i int) bool {
		if elem.row < l[i].row || elem.row == l[i].row && elem.val < l[i].val {
			return true
		}
		return false
	})

	if i == n {
		l = append(l, elem)
	} else {
		l = append(l[:i], append(List1{elem}, l[i:]...)...)
	}
	*list = l
}

func (l List1) Dump() []int {
	rst := []int{}
	for i := 0; i < len(l); i++ {
		rst = append(rst, l[i].val)
	}
	return rst
}

type Node struct {
	col int
	l   List1
}

type List []*Node

func (list *List) Insert(col, row, val int) {
	l := *list
	n := len(l)
	i := sort.Search(n, func(i int) bool {
		if col <= l[i].col {
			return true
		}
		return false
	})
	if i == n {
		l = append(l, &Node{col: col, l: List1{&Elem{row: row, val: val}}})
	} else if col == l[i].col {
		l[i].l.Insert(&Elem{row: row, val: val})
	} else {
		l = append(l[:i], append(List{&Node{col: col, l: List1{&Elem{row: row, val: val}}}}, l[i:]...)...)
	}
	*list = l
}

func (l List) Dump() [][]int {
	rst := [][]int{}
	for i := 0; i < len(l); i++ {
		rst = append(rst, l[i].l.Dump())
	}
	return rst
}
func verticalTraversal(root *TreeNode) [][]int {
	list := List{}

	var dfs func(node *TreeNode, col, row int)
	dfs = func(node *TreeNode, col, row int) {
		if node == nil {
			return
		}
		list.Insert(col, row, node.Val)

		dfs(node.Left, col-1, row+1)
		dfs(node.Right, col+1, row+1)
	}

	dfs(root, 0, 0)
	return list.Dump()
}

// 官方答案: 方法一样, 更加简洁
type data struct{ col, row, val int }

func verticalTraversal1(root *TreeNode) (ans [][]int) {
	nodes := []data{}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		nodes = append(nodes, data{col, row, node.Val})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]
		return a.col < b.col || a.col == b.col && (a.row < b.row || a.row == b.row && a.val < b.val)
	})

	lastCol := math.MinInt32
	for _, node := range nodes {
		if node.col != lastCol {
			lastCol = node.col
			ans = append(ans, nil)
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], node.val)
	}
	return
}
