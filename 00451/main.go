package main

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
)

type list []node

type node struct {
	c     rune
	count int
}

func (l *list) Len() int {
	return len(*l)
}
func (l *list) Less(i, j int) bool {
	return (*l)[i].count > (*l)[j].count
}
func (l *list) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func frequencySort(s string) string {
	l := make(list, 256)
	for i := 0; i < len(l); i++ {
		l[i].c = rune(i)
	}
	for _, c := range s {
		l[c].count += 1
	}
	sort.Sort(&l)
	b := bytes.NewBuffer(nil)
	w := bufio.NewWriter(b)
	for i := 0; i < len(l); i++ {
		if l[i].count > 0 {
			p := make([]byte, l[i].count)
			for j := 0; j < len(p); j++ {
				p[j] = byte(l[i].c)
			}
			w.Write(p)
		} else {
			break
		}
	}
	w.Flush()
	return b.String()
}

func main() {
	for _, test := range []struct {
		s      string
		expect string
	}{
		{"tree", "eetr"},
		{"cccaaa", "aaaccc"},
		{"Aabb", "bbAa"},
	} {
		r := frequencySort(test.s)
		fmt.Println(r, test.expect)
	}
}
