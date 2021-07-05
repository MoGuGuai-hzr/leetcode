package main

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func isNumber(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func s2i(s string) int {
	r := 0
	for i := range s {
		r = r*10 + int(s[i]) - '0'
	}
	return r
}

func isLowerCase(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}

type T int

const (
	number T = iota
	letter
	leftBracket
	rightBracket
)

func next(s string, i int) (string, int, T) {
	start := i
	c := s[i]
	var t T
	n := len(s)
	if isNumber(c) {
		// 数字
		t = number
		for isNumber(c) {
			i++
			if i >= n {
				break
			}
			c = s[i]
		}
	} else if c == '(' {
		// 左括号
		t = leftBracket
		i += 1
	} else if c == ')' {
		// 右括号
		t = rightBracket
		i += 1
	} else {
		// 字母
		t = letter

		for {
			i++
			if i >= n {
				break
			}
			c = s[i]
			if isLowerCase(c) {
				continue
			}
			break
		}
	}

	return s[start:i], i, t
}

type node struct {
	s string
	c int
}

type nodeList []node

func (n *nodeList) Len() int {
	return len(*n)
}

func (n *nodeList) Less(i int, j int) bool {
	s1 := (*n)[i]
	s2 := (*n)[j]
	return strings.Compare(s1.s, s2.s) < 0
}

func (n *nodeList) Swap(i int, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func (n nodeList) String() string {
	sort.Sort(&n)
	buf := new(bytes.Buffer)
	w := bufio.NewWriter(buf)
	for i := 0; i < len(n); i++ {
		node := n[i]
		w.WriteString(node.s)
		count := node.c
		for i+1 < len(n) && n[i+1].s == node.s {
			i++
			count += n[i].c
		}
		if count > 1 {
			w.WriteString(fmt.Sprintf("%d", count))
		}

	}
	w.Flush()
	return buf.String()
}

func countOfAtoms(formula string) string {
	m := make(map[int]nodeList)
	numBracket := 0
	m[numBracket] = []node{}

	now := node{}
	var s string
	var i int
	var t T

	for i < len(formula) {
		s, i, t = next(formula, i)
		switch t {
		case number:
			now.c = s2i(s)
			list := m[numBracket]
			list[len(list)-1] = now
			m[numBracket] = list

		case leftBracket:
			numBracket++
			m[numBracket] = []node{}

		case rightBracket:
			if i >= len(formula) {
				list := m[numBracket]
				numBracket--
				list2 := m[numBracket]
				for _, node := range list {
					list2 = append(list2, node)
				}
				m[numBracket] = list2
			} else {
				// 偷偷看一下下一个元素
				var tempI int
				s, tempI, t = next(formula, i)
				// 如果是数字, 全部扩大
				n := 1
				if t == number {
					n = s2i(s)
					i = tempI
				}
				list := m[numBracket]
				numBracket--
				list2 := m[numBracket]
				for _, node := range list {
					node.c *= n
					list2 = append(list2, node)
				}
				m[numBracket] = list2
			}
		default:
			now.s, now.c = s, 1
			list := m[numBracket]
			list = append(list, now)
			m[numBracket] = list
		}
	}

	list := m[numBracket]
	return list.String()
}

func main() {
	for _, test := range []struct {
		in, ex string
	}{
		{"(H)", "H"},
		{"H2O", "H2O"},
		{"Mg(OH)2", "H2MgO2"},
		{"Mg(OH)2Mg(OH)2", "H4Mg2O4"},
		{"Mg(H2O)N", "H2MgNO"},
	} {
		r := countOfAtoms(test.in)
		fmt.Println(r, test.ex, r == test.ex)
	}
}
